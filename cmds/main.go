package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9999"
)

func main() {

	// Open Listen on Port 9999
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Ada Error nih Om!! disini %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed Lagi Om!! disini %v", err.Error())
	}
}
