package main

import (
	"context"
	"database/sql"
	"log"
	pb "ecommerce-grpc/proto/pagamentos"
	pc "ecommerce-grpc/api-pagamentos/client"
)

type PagamentosHandler struct {
	pb.UnimplementedPagamentosServiceServer
	DB *sql.DB
}

func (h *PagamentosHandler) ProcessarPagamento(ctx context.Context, req *pb.ProcessarPagamentoRequest) (*pb.ProcessarPagamentoResponse, error) {
	log.Printf("[PAGAMENTOS] Recebendo pagamento do pedido: %d (valor: %.2f)", req.PedidoId, req.Valor)

	_, err := h.DB.Exec(
		`INSERT INTO pagamentos (pedido_id, data_processamento, status, metodo) VALUES (?, datetime('now'), ?, ?)`,
		req.PedidoId, "aprovado", "pix",
	)
	if err != nil {
		log.Printf("[PAGAMENTOS] Erro ao inserir pagamento: %v", err)
		return nil, err
	}

	log.Printf("[PAGAMENTOS] Pagamento registrado como 'aprovado' para pedido: %d", req.PedidoId)

	_, err = h.DB.Exec("UPDATE pedidos SET status=? WHERE id=?", "pagamento_aprovado", req.PedidoId)
	if err != nil {
		log.Printf("[PAGAMENTOS] Erro ao atualizar status do pedido: %v", err)
		return nil, err
	}

	log.Printf("[PAGAMENTOS] Status do pedido %d atualizado para 'pagamento_aprovado'", req.PedidoId)

	// Chama Estoque
	status, err := pc.SepararEstoqueGRPC(ctx, req.PedidoId)
	if err != nil {
		log.Printf("[PAGAMENTOS] Erro ao chamar separação de estoque para pedido %d: %v", req.PedidoId, err)
	} else {
		log.Printf("[PAGAMENTOS] Separação de estoque para pedido %d retornou status: %s", req.PedidoId, status)
	}

	return &pb.ProcessarPagamentoResponse{Status: status}, nil
}
