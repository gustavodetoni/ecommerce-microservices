package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "ecommerce-grpc/proto/fiscal"
	pc "ecommerce-grpc/fiscal/client"
)

type FiscalHandler struct {
	pb.UnimplementedFiscalServiceServer
	DB *sql.DB
}

func (h *FiscalHandler) EmitirNotaFiscal(ctx context.Context, req *pb.EmitirNotaFiscalRequest) (*pb.EmitirNotaFiscalResponse, error) {
	numero := fmt.Sprintf("NF-%d", req.PedidoId)
	data := "2025-05-29T22:55:00Z"
	chave := fmt.Sprintf("%d%d", req.PedidoId, 12345)

	log.Printf("[FISCAL] Emitindo nota fiscal para pedido: %d", req.PedidoId)

	res, err := h.DB.Exec(`INSERT INTO notas_fiscais (pedido_id, numero, data_emissao, chave_acesso) VALUES (?, ?, ?, ?)`,
		req.PedidoId, numero, data, chave)
	if err != nil {
		log.Printf("[FISCAL] Erro ao inserir nota fiscal: %v", err)
		return nil, err
	}
	nfID, _ := res.LastInsertId()
	log.Printf("[FISCAL] Nota fiscal emitida para pedido %d, NF-ID: %d", req.PedidoId, nfID)

	_, err = h.DB.Exec("UPDATE pedidos SET status=? WHERE id=?", "nf_emitida", req.PedidoId)
	if err != nil {
		log.Printf("[FISCAL] Erro ao atualizar status do pedido para 'nf_emitida': %v", err)
		return nil, err
	}
	log.Printf("[FISCAL] Status do pedido %d atualizado para 'nf_emitida'", req.PedidoId)

	// Chama Logística
	log.Printf("[FISCAL] Chamando serviço de logística para pedido %d, NF-ID: %d", req.PedidoId, nfID)
	status, _, err := pc.AgendarEnvioGRPC(ctx, req.PedidoId, nfID)
	if err != nil {
		log.Printf("[FISCAL] Erro ao agendar envio para pedido %d, NF-ID: %d: %v", req.PedidoId, nfID, err)
	} else {
		log.Printf("[FISCAL] Envio agendado, status: %s", status)
	}

	return &pb.EmitirNotaFiscalResponse{Status: status, NotaFiscalId: nfID}, nil
}
