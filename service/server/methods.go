package main

import (
	"context"
	"s0711-23/internal/demoapi"
)

// MySuperServer ...
type MySuperServer struct {
	demoapi.UnimplementedDemoServiceServer
}

// NewMySuperServer ...
func NewMySuperServer() *MySuperServer {
	return &MySuperServer{}
}

// GetUserByID ...
func (s *MySuperServer) GetUserByID(ctx context.Context, request *demoapi.GetUserRequest) (*demoapi.UserResponse, error) {
	response := &demoapi.UserResponse{
		Id:       request.Id,
		Username: "alex",
	}

	return response, nil
}
