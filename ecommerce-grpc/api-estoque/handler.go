package main

import (
	"context"
	"database/sql"
	"log"
	pb "ecommerce-grpc/proto/estoque"
	pc "ecommerce-grpc/api-estoque/client"
)

type EstoqueHandler struct {
	pb.UnimplementedEstoqueServiceServer
	DB *sql.DB
}

func (h *EstoqueHandler) SepararEstoque(ctx context.Context, req *pb.SepararEstoqueRequest) (*pb.SepararEstoqueResponse, error) {
	log.Printf("[ESTOQUE] Separando estoque para pedido: %d", req.PedidoId)

	_, err := h.DB.Exec("UPDATE pedidos SET status=? WHERE id=?", "em_separacao", req.PedidoId)
	if err != nil {
		log.Printf("[ESTOQUE] Erro ao atualizar status do pedido %d para 'em_separacao': %v", req.PedidoId, err)
		return nil, err
	}
	log.Printf("[ESTOQUE] Status do pedido %d atualizado para 'em_separacao'", req.PedidoId)

	// Chama Fiscal
	log.Printf("[ESTOQUE] Chamando serviço fiscal para emitir nota do pedido %d", req.PedidoId)
	status, err := pc.EmitirNotaFiscalGRPC(ctx, req.PedidoId)
	if err != nil {
		log.Printf("[ESTOQUE] Erro ao emitir nota fiscal para pedido %d: %v", req.PedidoId, err)
	} else {
		log.Printf("[ESTOQUE] Nota fiscal emitida para pedido %d, status retornado: %s", req.PedidoId, status)
	}

	return &pb.SepararEstoqueResponse{Status: status}, nil
}
