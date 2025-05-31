package main

import (
	"context"
	"database/sql"
	"log"
	"time"
	pb "ecommerce-grpc/proto/logistica"
)

type LogisticaHandler struct {
	pb.UnimplementedLogisticaServiceServer
	DB *sql.DB
}

func (h *LogisticaHandler) AgendarEnvio(ctx context.Context, req *pb.AgendarEnvioRequest) (*pb.AgendarEnvioResponse, error) {
	data := time.Now().Format(time.RFC3339)
	codigo := "Rastreio-" + data
	status := "enviado"

	log.Printf("[LOGISTICA] Agendando envio do pedido: %d, nota fiscal: %d", req.PedidoId, req.NotaFiscalId)

	_, err := h.DB.Exec(
		`INSERT INTO envios (pedido_id, nota_fiscal_id, data_despacho, codigo_rastreamento, status) VALUES (?, ?, ?, ?, ?)`,
		req.PedidoId, req.NotaFiscalId, data, codigo, status,
	)
	if err != nil {
		log.Printf("[LOGISTICA] Erro ao inserir envio: %v", err)
		return nil, err
	}

	_, err = h.DB.Exec("UPDATE pedidos SET status=? WHERE id=?", status, req.PedidoId)
	if err != nil {
		log.Printf("[LOGISTICA] Erro ao atualizar status do pedido: %v", err)
		return nil, err
	}

	log.Printf("[LOGISTICA] Envio agendado com sucesso para pedido: %d. Status: %s, Código de rastreamento: %s", req.PedidoId, status, codigo)

	return &pb.AgendarEnvioResponse{Status: status, CodigoRastreamento: codigo}, nil
}
