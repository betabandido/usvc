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
		count, err := client.GetCount(context.Background(), &v1.GetCountRequest{})
		if err != nil {
			log.Fatalf("error getting count: %v", err)
		}
		fmt.Printf("count: %v\n", count.Value)
	} else if *action == "update-count" {
		count, err := client.UpdateCount(context.Background(), &v1.UpdateCountRequest{Value: *value})
		if err != nil {
			log.Fatalf("error updating count: %v", err)
		}
		fmt.Printf("count: %v\n", count.Value)
	} else {
		log.Fatalf("unknown action: %v", action)
	}
}
