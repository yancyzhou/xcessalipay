package main

// main.go

import (
	"github.com/yancyzhou/xcessalipay/Config"
	"github.com/yancyzhou/xcessalipay/RPC"
	service "github.com/yancyzhou/xcessalipay/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// lis, err := net.Listen("tcp", ":"+Config.ServerConf.Server.Port)
	lis, err := net.Listen("tcp", ":"+Config.ServerConf.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}
	s := grpc.NewServer()
	service.RegisterAliPayServiceServer(s, &RPC.Myserver{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err.Error())
	}
}
