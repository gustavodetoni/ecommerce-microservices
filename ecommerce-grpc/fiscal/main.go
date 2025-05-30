package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	pb "ecommerce-grpc/proto/fiscal"
)

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
	if err != nil {
		log.Fatalf("Falha ao conectar banco: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFiscalServiceServer(grpcServer, &FiscalHandler{DB: db})
	log.Println("Fiscal rodando em :50054")
	grpcServer.Serve(lis)
}
