package main

import (
	"database/sql"
	pb "ecommerce-grpc/proto/pagamentos"
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
		log.Fatal(http.ListenAndServe(":2116", nil))
	}()

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
