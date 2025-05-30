package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	pb "ecommerce-grpc/proto/pedidos"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
	if err != nil {
		log.Fatalf("Falha ao conectar banco: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPedidosServiceServer(grpcServer, &PedidosHandler{DB: db})
	log.Println("Pedidos rodando em :50051")
	grpcServer.Serve(lis)
}
