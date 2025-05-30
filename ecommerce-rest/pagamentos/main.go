package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"encoding/json"
	"bytes"
	"net/http"
)

type PagamentoRequest struct {
	PedidoID int64   `json:"pedido_id"`
	Valor    float64 `json:"valor"`
}

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	router.POST("/pagamentos", func(c *gin.Context) {
		var p PagamentoRequest
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"erro": "Dados inválidos"})
			return
		}
		status := "aprovado"
		metodo := "pix"
		db.Exec(`INSERT INTO pagamentos (pedido_id, data_processamento, status, metodo) VALUES (?, ?, ?, ?)`,
			p.PedidoID, time.Now().Format(time.RFC3339), status, metodo)
		db.Exec("UPDATE pedidos SET status = ? WHERE id = ?", "pagamento_aprovado", p.PedidoID)

		payload := map[string]interface{}{"pedido_id": p.PedidoID}
		jsonData, _ := json.Marshal(payload)
		http.Post("http://localhost:8083/estoque/separar", "application/json", bytes.NewBuffer(jsonData))

		c.JSON(200, gin.H{"status": status})
	})

	router.Run(":8082")
}
