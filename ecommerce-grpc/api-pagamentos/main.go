package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	pb "ecommerce-grpc/proto/pagamentos"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
	if err != nil {
		log.Fatalf("Falha ao conectar banco: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPagamentosServiceServer(grpcServer, &PagamentosHandler{DB: db})
	log.Println("Pagamentos rodando em :50052")
	grpcServer.Serve(lis)
}
