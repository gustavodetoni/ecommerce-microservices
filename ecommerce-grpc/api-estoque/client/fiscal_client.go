package client

import (
	"context"
	"log"
	pb "ecommerce-grpc/proto/fiscal"
	"google.golang.org/grpc"
)

func EmitirNotaFiscalGRPC(ctx context.Context, pedidoID int64) (string, error) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Println("Erro ao conectar Fiscal:", err)
		return "erro", err
	}
	defer conn.Close()
	client := pb.NewFiscalServiceClient(conn)

	resp, err := client.EmitirNotaFiscal(ctx, &pb.EmitirNotaFiscalRequest{
		PedidoId: pedidoID,
	})
	if err != nil {
		return "erro", err
	}
	return resp.Status, nil
}
