package main

import (
    "log"
    "net"
    "database/sql"
    "google.golang.org/grpc"
    _ "github.com/mattn/go-sqlite3"
    pb "ecommerce-grpc/proto/catalogo"
)

func main() {
    lis, err := net.Listen("tcp", ":50056")
    if err != nil {
        log.Fatalf("Falha ao escutar: %v", err)
    }
    db, err := sql.Open("sqlite3", "../../database/ecommerce.db")
    if err != nil {
        log.Fatalf("Falha ao conectar banco: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterCatalogoServiceServer(grpcServer, &CatalogoHandler{DB: db})
    log.Println("Catalogo rodando em :50056")
    grpcServer.Serve(lis)
}
