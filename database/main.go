package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/brianvoe/gofakeit/v6"
)

var sqlCreateClientesTable = `
CREATE TABLE IF NOT EXISTS clientes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nome TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE
);`
var sqlCreateProdutosTable = `
CREATE TABLE IF NOT EXISTS produtos (
	id INTEGER PRIMARY KEY,
	nome TEXT NOT NULL,
	preco REAL NOT NULL,
	quantidade_estoque INTEGER NOT NULL
);`
var sqlCreatePedidosTable = `
CREATE TABLE IF NOT EXISTS pedidos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	cliente_id INTEGER NOT NULL,
	data_criacao TEXT NOT NULL,
	status TEXT NOT NULL,
	valor_total REAL,
	FOREIGN KEY (cliente_id) REFERENCES clientes (id)
);`
var sqlCreateItensPedidoTable = `
CREATE TABLE IF NOT EXISTS itens_pedido (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pedido_id INTEGER NOT NULL,
	produto_id INTEGER NOT NULL,
	quantidade INTEGER NOT NULL,
	preco_unitario REAL NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
	FOREIGN KEY (produto_id) REFERENCES produtos (id)
);`
var sqlCreatePagamentosTable = `
CREATE TABLE IF NOT EXISTS pagamentos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pedido_id INTEGER NOT NULL UNIQUE,
	data_processamento TEXT NOT NULL,
	status TEXT NOT NULL,
	metodo TEXT,
	FOREIGN KEY (pedido_id) REFERENCES pedidos (id)
);`
var sqlCreateNotasFiscaisTable = `
CREATE TABLE IF NOT EXISTS notas_fiscais (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pedido_id INTEGER NOT NULL UNIQUE,
	numero TEXT NOT NULL UNIQUE,
	data_emissao TEXT NOT NULL,
	chave_acesso TEXT,
	FOREIGN KEY (pedido_id) REFERENCES pedidos (id)
);`
var sqlCreateEnviosTable = `
CREATE TABLE IF NOT EXISTS envios (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	pedido_id INTEGER NOT NULL UNIQUE,
	nota_fiscal_id INTEGER NOT NULL UNIQUE,
	data_despacho TEXT,
	codigo_rastreamento TEXT,
	status TEXT NOT NULL,
	FOREIGN KEY (pedido_id) REFERENCES pedidos (id),
	FOREIGN KEY (nota_fiscal_id) REFERENCES notas_fiscais (id)
);`

func criarConexao(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func criarTabela(db *sql.DB, createTableSQL string) {
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}
}

func popularDados(db *sql.DB, numClientes, numProdutos, numPedidos int) {
	fmt.Println("Populando tabela clientes...")
	clientesIds := make([]int64, 0)
	for i := 0; i < numClientes; i++ {
		nome := gofakeit.Name()
		email := fmt.Sprintf("cliente_%d_%s@exemplo.com", i, gofakeit.Username())
		res, err := db.Exec("INSERT INTO clientes (nome, email) VALUES (?, ?)", nome, email)
		if err != nil {
			log.Fatalf("Erro ao inserir cliente: %v", err)
		}
		id, _ := res.LastInsertId()
		clientesIds = append(clientesIds, id)
	}

	fmt.Println("Populando tabela produtos...")
	produtosIds := make([]int64, 0)
	for i := 1; i <= numProdutos; i++ {
		nome := fmt.Sprintf("Produto %s %s", gofakeit.Word(), gofakeit.RandomString([]string{"Eletrônico", "Vestuário", "Livro", "Alimento"}))
		preco := gofakeit.Price(10.5, 500.99)
		estoque := gofakeit.Number(0, 100)
		_, err := db.Exec("INSERT INTO produtos (id, nome, preco, quantidade_estoque) VALUES (?, ?, ?, ?)", i, nome, preco, estoque)
		if err != nil {
			log.Fatalf("Erro ao inserir produto: %v", err)
		}
		produtosIds = append(produtosIds, int64(i))
	}

	fmt.Println("Populando tabelas pedidos e itens_pedido...")
	pedidosIdsCriados := make([]int64, 0)
	for i := 0; i < numPedidos; i++ {
		clienteId := clientesIds[rand.Intn(len(clientesIds))]
		dataCriacao := gofakeit.DateRange(time.Now().AddDate(0, 0, -90), time.Now()).Format(time.RFC3339)
		statusInicial := "pendente"
		res, err := db.Exec("INSERT INTO pedidos (cliente_id, data_criacao, status) VALUES (?, ?, ?)", clienteId, dataCriacao, statusInicial)
		if err != nil {
			log.Fatalf("Erro ao inserir pedido: %v", err)
		}
		pedidoId, _ := res.LastInsertId()
		pedidosIdsCriados = append(pedidosIdsCriados, pedidoId)

		numItens := gofakeit.Number(1, 5)
		valorTotalPedido := 0.0
		for j := 0; j < numItens; j++ {
			produtoId := produtosIds[rand.Intn(len(produtosIds))]
			row := db.QueryRow("SELECT preco FROM produtos WHERE id = ?", produtoId)
			var precoUnitario float64
			err := row.Scan(&precoUnitario)
			if err != nil {
				log.Printf("Produto ID %d não encontrado para o pedido %d", produtoId, pedidoId)
				continue
			}
			quantidade := gofakeit.Number(1, 3)
			valorTotalPedido += float64(quantidade) * precoUnitario
			_, err = db.Exec(`INSERT INTO itens_pedido (pedido_id, produto_id, quantidade, preco_unitario) VALUES (?, ?, ?, ?)`,
				pedidoId, produtoId, quantidade, precoUnitario)
			if err != nil {
				log.Fatalf("Erro ao inserir item_pedido: %v", err)
			}
		}
		_, err = db.Exec("UPDATE pedidos SET valor_total = ? WHERE id = ?", valorTotalPedido, pedidoId)
		if err != nil {
			log.Fatalf("Erro ao atualizar valor_total do pedido: %v", err)
		}
	}

	fmt.Println("Simulando status posteriores...")
	for _, pedidoId := range pedidosIdsCriados {
		if gofakeit.Float64Range(0, 1) < 0.8 {
			dataPagamento := gofakeit.DateRange(time.Now().AddDate(0, 0, -80), time.Now()).Format(time.RFC3339)
			statusPagamento := gofakeit.RandomString([]string{"aprovado", "rejeitado"})
			metodo := gofakeit.RandomString([]string{"cartao_credito", "boleto", "pix"})
			_, err := db.Exec(`INSERT INTO pagamentos (pedido_id, data_processamento, status, metodo) VALUES (?, ?, ?, ?)`,
				pedidoId, dataPagamento, statusPagamento, metodo)
			if err != nil {
				log.Fatalf("Erro ao inserir pagamento: %v", err)
			}
			if statusPagamento == "aprovado" {
				db.Exec("UPDATE pedidos SET status = 'pagamento_aprovado' WHERE id = ?", pedidoId)
				if gofakeit.Float64Range(0, 1) < 0.9 {
					db.Exec("UPDATE pedidos SET status = 'em_separacao' WHERE id = ?", pedidoId)
					dataNF := gofakeit.DateRange(time.Now().AddDate(0, 0, -70), time.Now()).Format(time.RFC3339)
					numeroNF := fmt.Sprintf("NF%d-%d", gofakeit.Number(10000, 99999), pedidoId)
					chaveNF := gofakeit.DigitN(44) + gofakeit.DigitN(44) + gofakeit.DigitN(44)
					res, err := db.Exec(`INSERT INTO notas_fiscais (pedido_id, numero, data_emissao, chave_acesso) VALUES (?, ?, ?, ?)`,
						pedidoId, numeroNF, dataNF, chaveNF)
					if err != nil {
						log.Fatalf("Erro ao inserir nota_fiscal: %v", err)
					}
					nfId, _ := res.LastInsertId()
					db.Exec("UPDATE pedidos SET status = 'nf_emitida' WHERE id = ?", pedidoId)
					if gofakeit.Float64Range(0, 1) < 0.95 {
						dataEnvio := gofakeit.DateRange(time.Now().AddDate(0, 0, -60), time.Now()).Format(time.RFC3339)
						rastreio := fmt.Sprintf("BR%dPY", gofakeit.Number(100000000, 999999999))
						statusEnvio := gofakeit.RandomString([]string{"aguardando_envio", "enviado", "entregue"})
						_, err := db.Exec(`INSERT INTO envios (pedido_id, nota_fiscal_id, data_despacho, codigo_rastreamento, status) VALUES (?, ?, ?, ?, ?)`,
							pedidoId, nfId, dataEnvio, rastreio, statusEnvio)
						if err != nil {
							log.Fatalf("Erro ao inserir envio: %v", err)
						}
						if statusEnvio == "enviado" || statusEnvio == "entregue" {
							db.Exec("UPDATE pedidos SET status = ? WHERE id = ?", statusEnvio, pedidoId)
						}
					}
				}
			}
		}
	}
	fmt.Println("População de dados concluída.")
}

func main() {
	database := "ecommerce.db"

	if _, err := os.Stat(database); err == nil {
		os.Remove(database)
		fmt.Printf("Banco de dados '%s' existente removido.\n", database)
	}

	db, err := criarConexao(database)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer db.Close()

	fmt.Println("Criando tabelas...")
	criarTabela(db, sqlCreateClientesTable)
	criarTabela(db, sqlCreateProdutosTable)
	criarTabela(db, sqlCreatePedidosTable)
	criarTabela(db, sqlCreateItensPedidoTable)
	criarTabela(db, sqlCreatePagamentosTable)
	criarTabela(db, sqlCreateNotasFiscaisTable)
	criarTabela(db, sqlCreateEnviosTable)
	fmt.Println("Tabelas criadas com sucesso.")

	popularDados(db, 50, 30, 100)

	fmt.Printf("Conexão com '%s' fechada.\n", database)
}
