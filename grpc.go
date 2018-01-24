package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

func startGRPC(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error when listening: %s", err.Error())
	}

	server := grpc.NewServer()
	RegisterServerServer(server, &entityServer{})
	server.Serve(lis)
}

type entityServer struct{}

// LookupEntity implements the Server
func (s *entityServer) LookupEntity(ctx context.Context, req *LookupRequest) (*LookupResponse, error) {
	return &LookupResponse{Entity: &Entity{Id: req.Id, Name: "Entity"}, Found: true}, nil
}

// SetEntity implements the Server
func (s *entityServer) SetEntity(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	if !validate(req) {
		return &SetResponse{
			Entity:  req.Entity,
			Reason:  "Unable to validate request",
			Success: false,
		}, nil
	}

	return &SetResponse{
		Entity:  req.Entity,
		Success: true,
	}, nil
}
