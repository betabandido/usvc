package main

import (
	"context"
	"flag"
	"fmt"
	v1 "github.com/betabandido/usvc/proto/v1"
	"google.golang.org/grpc"
	"log"
)

var port = flag.Int("port", 5000, "port number")
var action = flag.String("action", "get-count", "action")
var value = flag.Int64("value", 1, "value")
var addDelay = flag.Bool("add-delay", false, "add random delay when getting count")
var errorRate = flag.Int("error-rate", 0, "error rate [0..100]")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", *port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("error connecting to server: %v", err)
	}
	defer conn.Close()

	client := v1.NewCountServiceClient(conn)

	if *action == "get-count" {
		fmt.Printf("count: %v\n", getCount(client))
	} else if *action == "update-count" {
		fmt.Printf("count: %v\n", updateCount(client, *value))
	} else {
		log.Fatalf("unknown action: %v", action)
	}
}

func getCount(client v1.CountServiceClient) int64 {
	count, err := client.GetCount(context.Background(), &v1.GetCountRequest{
		AddDelay: *addDelay,
		ErrorRate: int32(*errorRate),
	})
	if err != nil {
		log.Fatalf("error getting count: %v", err)
	}

	return count.Value
}

func updateCount(client v1.CountServiceClient, value int64) int64 {
	count, err := client.UpdateCount(context.Background(), &v1.UpdateCountRequest{Value: value})
	if err != nil {
		log.Fatalf("error updating count: %v", err)
	}

	return count.Value
}
