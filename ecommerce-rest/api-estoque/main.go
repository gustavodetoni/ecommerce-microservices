package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"encoding/json"
	"bytes"
	"net/http"
	"github.com/zsais/go-gin-prometheus"
)

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	p:= ginprometheus.NewPrometheus("api-estoque")
	p.Use(router)

	router.POST("/estoque/separar", func(c *gin.Context) {
		var req struct {
			PedidoID int64 `json:"pedido_id"`
		}
		c.ShouldBindJSON(&req)
		db.Exec("UPDATE pedidos SET status=? WHERE id=?", "em_separacao", req.PedidoID)

		payload := map[string]interface{}{"pedido_id": req.PedidoID}
		jsonData, _ := json.Marshal(payload)
		http.Post("http://localhost:8084/fiscal/emitir", "application/json", bytes.NewBuffer(jsonData))
		c.JSON(200, gin.H{"status": "separado"})
	})

	router.Run(":8083")
}
