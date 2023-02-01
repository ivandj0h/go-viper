package services

import (
	"context"
	pagingPb "go-grpc/pb/pagination"
	productPb "go-grpc/pb/product"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
}

func (p *ProductService) GetAllProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{
		Pagination: &pagingPb.Pagination{
			Total:       5,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    2,
		},
		Data: []*productPb.Product{
			{
				Id:    1,
				Name:  "Celana Dalam",
				Price: 150000,
				Stock: 10,
				Category: &productPb.Category{
					Id:   1,
					Name: "Celana",
				},
			},
			{
				Id:    2,
				Name:  "Sempak Ungu",
				Price: 150000,
				Stock: 10,
				Category: &productPb.Category{
					Id:   1,
					Name: "Celana",
				},
			},
		},
	}

	return products, nil
}
