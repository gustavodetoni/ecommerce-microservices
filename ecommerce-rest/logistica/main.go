package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	router.POST("/logistica/enviar", func(c *gin.Context) {
		var req struct {
			PedidoID     int64 `json:"pedido_id"`
			NotaFiscalID int64 `json:"nota_fiscal_id"`
		}
		c.ShouldBindJSON(&req)
		data := time.Now().Format(time.RFC3339)
		codigo := "Rastreio-" + data
		status := "enviado"
		db.Exec(`INSERT INTO envios (pedido_id, nota_fiscal_id, data_despacho, codigo_rastreamento, status) VALUES (?, ?, ?, ?, ?)`,
			req.PedidoID, req.NotaFiscalID, data, codigo, status)
		db.Exec("UPDATE pedidos SET status=? WHERE id=?", "enviado", req.PedidoID)
		c.JSON(200, gin.H{"status": "enviado", "rastreio": codigo})
	})

	router.Run(":8085")
}
