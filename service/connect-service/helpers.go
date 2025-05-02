package main

import "s0711-23/internal/proxyproto"

func respondError(code uint32, msg string) *proxyproto.ConnectResponse {
	return &proxyproto.ConnectResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: msg,
		},
	}
}
