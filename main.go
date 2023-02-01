package main

import (
	"fmt"
	pb "go-grpc/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Celana Dalam",
				Price: 50000.00,
				Stock: 2,
				Category: &pb.Category{
					Id:   1,
					Name: "celana",
				},
			},
			{
				Id:    2,
				Name:  "Celana Pendek",
				Price: 150000.00,
				Stock: 22,
				Category: &pb.Category{
					Id:   1,
					Name: "celana",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Error Om!", err)
	}
	fmt.Println(data)
}
