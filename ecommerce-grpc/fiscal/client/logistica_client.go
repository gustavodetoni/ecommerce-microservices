package client

import (
	"context"
	"log"
	pb "ecommerce-grpc/proto/logistica"
	"google.golang.org/grpc"
)

func AgendarEnvioGRPC(ctx context.Context, pedidoID, notaFiscalID int64) (string, string, error) {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Println("Erro ao conectar Logistica:", err)
		return "erro", "", err
	}
	defer conn.Close()
	client := pb.NewLogisticaServiceClient(conn)

	resp, err := client.AgendarEnvio(ctx, &pb.AgendarEnvioRequest{
		PedidoId:     pedidoID,
		NotaFiscalId: notaFiscalID,
	})
	if err != nil {
		return "erro", "", err
	}
	return resp.Status, resp.CodigoRastreamento, nil
}
