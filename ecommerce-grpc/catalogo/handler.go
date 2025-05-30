package main

import (
    "context"
    "database/sql"
    "log"
    pb "ecommerce-grpc/proto/catalogo"
)

type CatalogoHandler struct {
    pb.UnimplementedCatalogoServiceServer
    DB *sql.DB
}

func (h *CatalogoHandler) ListarProdutos(ctx context.Context, req *pb.ListarProdutosRequest) (*pb.ListarProdutosResponse, error) {
    log.Println("[CATALOGO] Recebida solicitação para listar produtos.")

    rows, err := h.DB.Query("SELECT id, nome, quantidade_estoque FROM produtos")
    if err != nil {
        log.Printf("[CATALOGO] Erro ao buscar produtos: %v", err)
        return nil, err
    }
    defer rows.Close()

    produtos := []*pb.Produto{}
    for rows.Next() {
        var id int64
        var nome string
        var estoque int32
        if err := rows.Scan(&id, &nome, &estoque); err != nil {
            log.Printf("[CATALOGO] Erro ao ler produto do banco: %v", err)
            return nil, err
        }
        produtos = append(produtos, &pb.Produto{
            Id:      id,
            Nome:    nome,
            Estoque: estoque,
        })
    }
    log.Printf("[CATALOGO] %d produtos retornados.", len(produtos))
    return &pb.ListarProdutosResponse{Produtos: produtos}, nil
}
