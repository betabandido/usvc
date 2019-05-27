package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/betabandido/usvc/proto/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var port = flag.Int("port", 5000, "port number")
var children = flag.String("children", "", "child services")

type countServer struct {
	childConnections []*grpc.ClientConn
	childCountServices []v1.CountServiceClient
	value              int64
}

func makeCountServer(children []int, tracer opentracing.Tracer) *countServer {
	childConnections := make([]*grpc.ClientConn, 0)
	for _, port := range children {
		conn, err := grpc.Dial(
			fmt.Sprintf("localhost:%d", port),
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
				grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
			)),
		)
		if err != nil {
			log.Fatalf("error creating connection: %v", err)
		}

		childConnections = append(childConnections, conn)
	}

	childCountServices := make([]v1.CountServiceClient, 0)
	for _, conn := range childConnections {
		childCountServices = append(childCountServices, v1.NewCountServiceClient(conn))
	}

	return &countServer{
		childConnections: childConnections,
		childCountServices: childCountServices,
		value: 0,
	}
}

func (s *countServer) tearDown() {
	for _, conn := range s.childConnections {
		err := conn.Close()
		if err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}
}

func (s *countServer) GetCount(ctx context.Context, request *v1.GetCountRequest) (*v1.Count, error) {
	delay := 0
	if request.AddDelay {
		delay = rand.Intn(2000) + 250
	}

	log.Printf("Processing GetCount with delay %d", delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)

	totalValue := s.value

	for _, child := range s.childCountServices {
		count, err := child.GetCount(ctx, &v1.GetCountRequest{
			AddDelay: request.AddDelay,
		})
		if err != nil {
			log.Printf("error getting count: %v", err)
		} else {
			totalValue += count.Value
		}
	}

	return &v1.Count{Value: totalValue}, nil
}

func (s *countServer) UpdateCount(ctx context.Context, request *v1.UpdateCountRequest) (*v1.Count, error) {
	s.value = request.Value
	return &v1.Count{Value: s.value}, nil
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer, closer, err := newTracer()
	defer closer.Close()
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)

	countServer := makeCountServer(childPorts(*children), tracer)
	defer countServer.tearDown()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
	)
	v1.RegisterCountServiceServer(server, countServer)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("error serving requests: %v", err)
	}
}

func newTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			BufferFlushInterval: 1 * time.Second,
		},
		ServiceName: serviceName(),
	}

	return cfg.NewTracer()
}

func serviceName() string {
	return fmt.Sprintf("usvc-%d", *port)
}

func childPorts(str string) []int {
	ports := make([]int, 0)
	if len(str) == 0 {
		return ports
	}

	for _, elem := range strings.Split(str, ",") {
		port, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatalf("error parsing port: %v", err)
		}
		ports = append(ports, port)
	}
	return ports
}
