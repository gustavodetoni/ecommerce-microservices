package main

import (
	"database/sql"
	pb "ecommerce-grpc/proto/logistica"
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
		log.Fatal(http.ListenAndServe(":2115", nil))
	}()

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
