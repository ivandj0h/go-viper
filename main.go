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
	// it became compact binary wire format
	/**
	 *
	 * [18 39 8 1 18 12 67 101 108 97 110 97 32 68 97 108 97 109 25 0 0 0 0 0 106 232 64 32 2 42 10 8 1 18 6 99 101 108 97 110 97 18 40 8 2 18 13 67 101 108 97 110 97 32 80 101 110 100 101 107 25 0 0 0 0 128 79 2 65 32 22 42 10 8 1 18 6 99 101 108 97 110 97]
	 *
	 */
	fmt.Println(data)

	fmt.Println("-------- UnMarshal --------")

	// decode using UnMarshal
	hasilProduct := &pb.Products{}
	if err = proto.Unmarshal(data, hasilProduct); err != nil {
		log.Fatal("Unsuccess!", err)
	}

	fmt.Println(hasilProduct)

	// loop data inside product
	for _, product := range hasilProduct.GetData() {
		fmt.Println(product.GetName())
	}
}
