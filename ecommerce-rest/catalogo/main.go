package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "../../database/ecommerce.db")
	defer db.Close()
	router := gin.Default()

	router.GET("/catalogo", func(c *gin.Context) {
		rows, _ := db.Query("SELECT id, nome, quantidade_estoque FROM produtos")
		defer rows.Close()
		produtos := []map[string]interface{}{}
		for rows.Next() {
			var id, estoque int
			var nome string
			rows.Scan(&id, &nome, &estoque)
			produtos = append(produtos, map[string]interface{}{
				"id": id, "nome": nome, "estoque": estoque,
			})
		}
		c.JSON(200, produtos)
	})

	router.Run(":8086")
}
