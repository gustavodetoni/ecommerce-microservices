package main

import (
	"context"
	"database/sql"
	"log"
	pb "ecommerce-grpc/proto/pedidos"
	pc "ecommerce-grpc/api-pedidos/client"
)

type PedidosHandler struct {
	pb.UnimplementedPedidosServiceServer
	DB *sql.DB
}

func (h *PedidosHandler) CriarPedido(ctx context.Context, req *pb.CriarPedidoRequest) (*pb.CriarPedidoResponse, error) {
	log.Printf("[PEDIDOS] Recebendo novo pedido para cliente: %d", req.ClienteId)

	tx, err := h.DB.Begin()
	if err != nil {
		log.Printf("[PEDIDOS] Erro ao iniciar transação: %v", err)
		return nil, err
	}

	res, err := tx.Exec(`INSERT INTO pedidos (cliente_id, data_criacao, status) VALUES (?, datetime('now'), ?)`, req.ClienteId, "pendente")
	if err != nil {
		tx.Rollback()
		log.Printf("[PEDIDOS] Erro ao inserir pedido: %v", err)
		return nil, err
	}
	pedidoID, _ := res.LastInsertId()
	log.Printf("[PEDIDOS] Pedido criado com id: %d", pedidoID)

	valorTotal := 0.0
	for _, item := range req.Itens {
		var preco float64
		h.DB.QueryRow("SELECT preco FROM produtos WHERE id = ?", item.ProdutoId).Scan(&preco)
		valorTotal += preco * float64(item.Quantidade)
		_, err := tx.Exec(
			`INSERT INTO itens_pedido (pedido_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?)`,
			pedidoID, item.ProdutoId, item.Quantidade, preco)
		if err != nil {
			tx.Rollback()
			log.Printf("[PEDIDOS] Erro ao inserir item do pedido: %v", err)
			return nil, err
		}
		log.Printf("[PEDIDOS] Item adicionado ao pedido %d: produto %d, quantidade %d, preco %.2f",
			pedidoID, item.ProdutoId, item.Quantidade, preco)
	}

	_, err = tx.Exec("UPDATE pedidos SET valor_total=? WHERE id=?", valorTotal, pedidoID)
	if err != nil {
		tx.Rollback()
		log.Printf("[PEDIDOS] Erro ao atualizar valor_total do pedido: %v", err)
		return nil, err
	}
	tx.Commit()
	log.Printf("[PEDIDOS] Valor total do pedido %d: %.2f", pedidoID, valorTotal)

	// Chama Pagamentos
	log.Printf("[PEDIDOS] Chamando serviço de pagamentos para pedido %d", pedidoID)
	status, err := pc.ProcessarPagamentoGRPC(ctx, pedidoID, valorTotal)
	if err != nil {
		log.Printf("[PEDIDOS] Erro ao processar pagamento do pedido %d: %v", pedidoID, err)
	} else {
		log.Printf("[PEDIDOS] Pagamento processado, status retornado: %s", status)
	}

	return &pb.CriarPedidoResponse{
		PedidoId: pedidoID,
		Status:   status,
	}, nil
}
