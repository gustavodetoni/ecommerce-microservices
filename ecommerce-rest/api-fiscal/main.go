package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"github.com/zsais/go-gin-prometheus"
)

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	p:= ginprometheus.NewPrometheus("api-fiscal")
	p.Use(router)

	router.POST("/fiscal/emitir", func(c *gin.Context) {
		var req struct {
			PedidoID int64 `json:"pedido_id"`
		}
		c.ShouldBindJSON(&req)
		numero := fmt.Sprintf("NF-%d", req.PedidoID)
		data := time.Now().Format(time.RFC3339)
		chave := fmt.Sprintf("%d%d", req.PedidoID, time.Now().UnixNano())
		res, _ := db.Exec(`INSERT INTO notas_fiscais (pedido_id, numero, data_emissao, chave_acesso) VALUES (?, ?, ?, ?)`,
			req.PedidoID, numero, data, chave)
		nfID, _ := res.LastInsertId()
		db.Exec("UPDATE pedidos SET status=? WHERE id=?", "nf_emitida", req.PedidoID)

		payload := map[string]interface{}{"pedido_id": req.PedidoID, "nota_fiscal_id": nfID}
		jsonData, _ := json.Marshal(payload)
		http.Post("http://localhost:8085/logistica/enviar", "application/json", bytes.NewBuffer(jsonData))
		c.JSON(200, gin.H{"status": "nota_emitida"})
	})

	router.Run(":8084")
}
