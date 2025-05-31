package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	url         = "http://localhost:8081/pedidos" 
	concurrent  = 50   
	total       = 1000
)

type Pedido struct {
	ClienteID int `json:"cliente_id"`
	Itens     []struct {
		ProdutoID int `json:"produto_id"`
		Quantidade int `json:"quantidade"`
	} `json:"itens"`
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	var success, fail int
	var totalTime time.Duration

	fmt.Printf("Iniciando carga REST: %d pedidos, %d concorrentes\n", total, concurrent)
	sem := make(chan struct{}, concurrent)

	for i := 0; i < total; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(n int) {
			defer wg.Done()
			pedido := map[string]interface{}{
				"cliente_id":  1,
				"itens":       []map[string]interface{}{{"produto_id": 1, "quantidade": 1}},
			}
			payload, _ := json.Marshal(pedido)
			t1 := time.Now()
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
			latency := time.Since(t1)
			mu.Lock()
			totalTime += latency
			if err != nil || resp.StatusCode >= 400 {
				fail++
			} else {
				success++
			}
			mu.Unlock()
			if resp != nil {
				resp.Body.Close()
			}
			<-sem
		}(i)
	}
	wg.Wait()
	dur := time.Since(start)
	fmt.Printf("REST FINALIZADO\nTotal: %d | Sucesso: %d | Falhas: %d\n", total, success, fail)
	fmt.Printf("Latência média: %.2f ms | Throughput: %.2f req/s\n",
		(float64(totalTime.Milliseconds())/float64(total)), float64(total)/dur.Seconds())
}
