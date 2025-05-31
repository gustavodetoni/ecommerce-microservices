package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "ecommerce-grpc/proto/pedidos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr        = "localhost:50051"
	concurrent  = 50
	total       = 1000
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	var success, fail int
	var totalTime time.Duration

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewPedidosServiceClient(conn)

	fmt.Printf("Iniciando carga gRPC: %d pedidos, %d concorrentes\n", total, concurrent)
	sem := make(chan struct{}, concurrent)

	for i := 0; i < total; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(n int) {
			defer wg.Done()
			req := &pb.CriarPedidoRequest{
				ClienteId: 1,
				Itens: []*pb.ItemPedido{
					{ProdutoId: 1, Quantidade: 1},
				},
			}
			t1 := time.Now()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			_, err := client.CriarPedido(ctx, req)
			latency := time.Since(t1)
			mu.Lock()
			totalTime += latency
			if err != nil {
				fail++
			} else {
				success++
			}
			mu.Unlock()
			<-sem
		}(i)
	}
	wg.Wait()
	dur := time.Since(start)
	fmt.Printf("gRPC FINALIZADO\nTotal: %d | Sucesso: %d | Falhas: %d\n", total, success, fail)
	fmt.Printf("Latência média: %.2f ms | Throughput: %.2f req/s\n",
		(float64(totalTime.Milliseconds())/float64(total)), float64(total)/dur.Seconds())
}
