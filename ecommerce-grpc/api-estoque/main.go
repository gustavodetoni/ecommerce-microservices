package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	pb "ecommerce-grpc/proto/estoque"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
	if err != nil {
		log.Fatalf("Falha ao conectar banco: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEstoqueServiceServer(grpcServer, &EstoqueHandler{DB: db})
	log.Println("Estoque rodando em :50053")
	grpcServer.Serve(lis)
}
