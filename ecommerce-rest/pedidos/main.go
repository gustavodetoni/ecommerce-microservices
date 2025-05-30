package main

import (
	"database/sql"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"bytes"
	"encoding/json"
)

type NovoPedido struct {
	ClienteID int              `json:"cliente_id"`
	Items     []ItemNovoPedido `json:"itens"`
}
type ItemNovoPedido struct {
	ProdutoID int `json:"produto_id"`
	Quantidade int `json:"quantidade"`
}

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	router.POST("/pedidos", func(c *gin.Context) {
		var np NovoPedido
		if err := c.ShouldBindJSON(&np); err != nil {
			c.JSON(400, gin.H{"erro": "Dados inválidos"})
			return
		}
		tx, _ := db.Begin()
		res, err := tx.Exec(`INSERT INTO pedidos (cliente_id, data_criacao, status) VALUES (?, ?, ?)`,
			np.ClienteID, time.Now().Format(time.RFC3339), "pendente")
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"erro": "Falha ao inserir pedido"})
			return
		}
		pedidoID, _ := res.LastInsertId()
		valorTotal := 0.0
		for _, item := range np.Items {
			var preco float64
			db.QueryRow("SELECT preco FROM produtos WHERE id = ?", item.ProdutoID).Scan(&preco)
			valorTotal += preco * float64(item.Quantidade)
			tx.Exec(`INSERT INTO itens_pedido (pedido_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?)`,
				pedidoID, item.ProdutoID, item.Quantidade, preco)
		}
		tx.Exec("UPDATE pedidos SET valor_total=? WHERE id=?", valorTotal, pedidoID)
		tx.Commit()

		pagReq := map[string]interface{}{
			"pedido_id": pedidoID,
			"valor":     valorTotal,
		}
		jsonPag, _ := json.Marshal(pagReq)
		http.Post("http://localhost:8082/pagamentos", "application/json", bytes.NewBuffer(jsonPag))

		c.JSON(201, gin.H{"pedido_id": pedidoID, "status": "pendente"})
	})

	router.Run(":8081")
}
