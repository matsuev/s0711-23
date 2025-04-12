package main

import (
	"context"
	"encoding/json"
	"s0711-23/internal/proxyproto"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
}

// Connect ...
func (s *Service) Connect(ctx context.Context, request *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {

	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	authRequest := &AuthRequest{}

	if err := json.Unmarshal(request.Data, authRequest); err != nil {
		return nil, err
	}

	if authRequest.Username != "root" || authRequest.Password != "secret" {
		return &proxyproto.ConnectResponse{
			Error: &proxyproto.Error{
				Code:    101,
				Message: "unauthorized",
			},
		}, nil
	}

	response := &proxyproto.ConnectResponse{
		Result: &proxyproto.ConnectResult{
			User: "root",
		},
	}

	return response, nil
}
