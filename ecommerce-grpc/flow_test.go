package main

import (
    "context"
    "log"
    "testing"
    "time"

    catpb "ecommerce-grpc/proto/catalogo"
    pedpb "ecommerce-grpc/proto/pedidos"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func TestFluxoCompleto(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    connCat, err := grpc.Dial("localhost:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Erro ao conectar ao serviço catalogo: %v", err)
    }
    defer connCat.Close()
    catClient := catpb.NewCatalogoServiceClient(connCat)
    catResp, err := catClient.ListarProdutos(ctx, &catpb.ListarProdutosRequest{})
    if err != nil {
        t.Fatalf("Erro ao listar produtos: %v", err)
    }
    if len(catResp.Produtos) == 0 {
        t.Fatalf("Nenhum produto disponível para testar")
    }
    produto := catResp.Produtos[0]

    connPed, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("Erro ao conectar ao serviço pedidos: %v", err)
    }
    defer connPed.Close()
    pedClient := pedpb.NewPedidosServiceClient(connPed)
    pedResp, err := pedClient.CriarPedido(ctx, &pedpb.CriarPedidoRequest{
        ClienteId:  1,
        Itens: []*pedpb.ItemPedido{
            {
                ProdutoId: produto.Id,
                Quantidade: 1,
            },
        },
    })
    if err != nil {
        t.Fatalf("Erro ao criar pedido: %v", err)
    }
    if pedResp.PedidoId == 0 {
        t.Fatalf("PedidoId retornado é inválido")
    }
    log.Printf("Pedido criado com id: %d e status: %s", pedResp.PedidoId, pedResp.Status)
}
