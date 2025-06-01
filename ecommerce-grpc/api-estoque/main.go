package main

import (
	"database/sql"
	pb "ecommerce-grpc/proto/estoque"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":2113", nil))
	}()

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
