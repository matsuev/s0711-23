package main

import (
	"context"
	"encoding/json"
	"log"
	"s0711-23/internal/proxyproto"
	"s0711-23/internal/userdb"
	"strconv"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	query userdb.Querier
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

	result, err := s.query.UserLogin(context.Background(), userdb.UserLoginParams{
		Username: authRequest.Username,
		Password: authRequest.Password,
	})

	if err != nil {
		log.Println(err)
		return &proxyproto.ConnectResponse{
			Error: &proxyproto.Error{
				Code:    101,
				Message: "unauthorized",
			},
		}, nil
	}

	if result.Enabled {
		return &proxyproto.ConnectResponse{
			Result: &proxyproto.ConnectResult{
				User: strconv.Itoa(int(result.ID)),
			},
		}, nil
	}

	return &proxyproto.ConnectResponse{
		Error: &proxyproto.Error{
			Code:    101,
			Message: "unauthorized",
		},
	}, nil
}
