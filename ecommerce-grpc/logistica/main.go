package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	pb "ecommerce-grpc/proto/logistica"
)

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
	if err != nil {
		log.Fatalf("Falha ao conectar banco: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLogisticaServiceServer(grpcServer, &LogisticaHandler{DB: db})
	log.Println("Logistica rodando em :50055")
	grpcServer.Serve(lis)
}
