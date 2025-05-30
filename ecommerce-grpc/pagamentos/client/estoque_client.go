package client

import (
	"context"
	"log"
	pb "ecommerce-grpc/proto/estoque"
	"google.golang.org/grpc"
)

func SepararEstoqueGRPC(ctx context.Context, pedidoID int64) (string, error) {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Println("Erro ao conectar Estoque:", err)
		return "erro", err
	}
	defer conn.Close()
	client := pb.NewEstoqueServiceClient(conn)

	resp, err := client.SepararEstoque(ctx, &pb.SepararEstoqueRequest{
		PedidoId: pedidoID,
	})
	if err != nil {
		return "erro", err
	}
	return resp.Status, nil
}
