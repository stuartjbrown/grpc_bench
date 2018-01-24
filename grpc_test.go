package main

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
)

// BenchmarkGRPCLookup runs a benchmark test of the GRPC LookupEntity endpoint
func BenchmarkGRPCLookup(b *testing.B) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		_, err := client.LookupEntity(context.Background(), &LookupRequest{Id: 1234})
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGRPCSetEntity runs a benchmark test of the GRPC SetEntity endpoint
func BenchmarkGRPCSetEntity(b *testing.B) {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewServerClient(conn)

	// run grpc calls against it
	for i := 0; i < b.N; i++ {
		_, err := client.SetEntity(context.Background(), &SetRequest{Entity: &Entity{Id: 1234, Name: "SomeEntity"}})
		if err != nil {
			b.Fatal(err)
		}
	}
}
