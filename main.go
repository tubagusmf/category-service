package main

import (
	"log"
	"net"

	grpcSvc "github.com/tubagusmf/category-service/internal/delivery/grpc"
	"github.com/tubagusmf/category-service/internal/repository"
	"github.com/tubagusmf/category-service/internal/usecase"
	pb "github.com/tubagusmf/category-service/pb/category"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	categoryRepo := repository.NewCategoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	categoryService := grpcSvc.NewCategoryService(categoryUsecase)

	pb.RegisterCategoryServiceServer(s, categoryService)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting gRPC server on :3200")

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
