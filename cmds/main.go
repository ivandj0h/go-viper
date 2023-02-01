package main

import (
	"log"
	"net"

	"go-grpc/cmds/services"
	productPb "go-grpc/pb/product"

	"google.golang.org/grpc"
)

const (
	port = ":7000"
)

func main() {

	// Open Listen on Port 9999
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Ada Error nih Om!! disini %v", err.Error())
	}

	grpcServer := grpc.NewServer()

	productService := services.ProductService{}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server Starting on %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed Lagi Om!! disini %v", err.Error())
	}
}
