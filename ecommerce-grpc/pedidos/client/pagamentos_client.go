package client

import (
	"context"
	"log"
	pb "ecommerce-grpc/proto/pagamentos"
	"google.golang.org/grpc"
)

func ProcessarPagamentoGRPC(ctx context.Context, pedidoID int64, valor float64) (string, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Println("Erro ao conectar Pagamentos:", err)
		return "erro", err
	}
	defer conn.Close()
	client := pb.NewPagamentosServiceClient(conn)

	resp, err := client.ProcessarPagamento(ctx, &pb.ProcessarPagamentoRequest{
		PedidoId: pedidoID,
		Valor:    valor,
	})
	if err != nil {
		return "erro", err
	}
	return resp.Status, nil
}
