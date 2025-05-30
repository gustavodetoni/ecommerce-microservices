# Checkout Microservices

Projeto de checkout distribuído em microserviços REST e GRPC  usando Golang e SQLite.

## Visão Geral

O sistema simula o fluxo de um pedido de e-commerce, dividido em 6 serviços independentes, que se comunicam via HTTP REST e GRPC.  
Cada serviço tem responsabilidade única e utiliza a mesma base SQLite localizada em `database/ecommerce.db`.

---

## Serviços e Responsabilidades

- **pedidos**  
  Responsável por criar pedidos, salvar itens do pedido e iniciar o fluxo de checkout.
- **pagamentos**  
  Simula o processamento de pagamento e avança o fluxo caso aprovado.
- **estoque**  
  Simula a separação dos itens no estoque para o pedido.
- **fiscal**  
  Simula a emissão da nota fiscal após a separação do estoque.
- **logistica**  
  Simula o envio do pedido para a transportadora e finaliza o fluxo.
- **catalogo**  
  Exibe o catálogo de produtos em estoque.

---

# Rotas para teste do Ecommerce-rest

## 1. Serviço de Catálogo

### GET /catalogo
Lista todos os produtos em estoque.

**Retorno:**
```json
[
  {
    "id": 1,
    "nome": "Produto X",
    "estoque": 50
  }
]
```

## 2. Serviço de Pedidos

### POST /pedidos
Cria um novo pedido e inicia o fluxo.

**Corpo (raw JSON:**
```json
{
  "cliente_id": 1,
  "itens": [
    {
      "produto_id": 1,
      "quantidade": 2
    },
    {
      "produto_id": 2,
      "quantidade": 1
    }
  ]
}
```

**Retorno:**
```json
{
  "pedido_id": 101,
  "status": "pendente"
}
```

## 3. Serviço de Pagamentos

### POST /pagamentos
Usado internamente pelo serviço de pedidos. Pode ser testado diretamente para simular pagamentos.

**Corpo (raw JSON):**
```json
{
  "pedido_id": 101,
  "valor": 120.00
}
```

**Retorno:**
```json
{
  "status": "aprovado"
}
```

## 4. Serviço de Estoque

### POST /estoque/separar
Usado internamente pelo serviço de pagamentos. Simula a separação dos itens do pedido.

**Corpo (raw JSON):**
```json
{
  "pedido_id": 101
}
```

**Retorno:**
```json
{
  "status": "separado"
}
```

## 5. Serviço Fiscal

### POST /fiscal/emitir
Usado internamente pelo serviço de estoque. Simula emissão de nota fiscal.

**Corpo (raw JSON):**
```json
{
  "pedido_id": 101
}
```

**Retorno:**
```json
{
  "status": "nota_emitida"
}
```

## 6. Serviço de Logística

### POST /logistica/enviar
Usado internamente pelo serviço fiscal. Simula envio do pedido.

**Corpo (raw JSON):**
```json
{
  "pedido_id": 101,
  "nota_fiscal_id": 77
}
```

**Retorno:**
```json
{
  "status": "enviado",
  "rastreio": "Rastreio-2025-05-29T22:55:00Z"
}
```


---

# Rotas para teste do Ecommerce-grpc

Já existe um teste chamado flow_test.go. Para rodar:
```
go test -v flow_test.go

```
Esse teste executa o fluxo:

1. Consulta produtos no catálogo
2. Cria um novo pedido
3. Todo o restante do fluxo acontece automaticamente (pagamento, estoque, fiscal, logística)

Verifique os terminais dos serviços para os logs detalhados do fluxo.

## Fluxo dos Serviços

1. **Catálogo** - Consulta produtos disponíveis
2. **Pedidos** - Inicia o processo criando um pedido
3. **Pagamentos** - Processa o pagamento do pedido
4. **Estoque** - Separa os itens do pedido
5. **Fiscal** - Emite a nota fiscal
6. **Logística** - Realiza o envio e gera código de rastreamento

## Portas dos Serviços

- Pedidos: `8081`
- Pagamentos: `8082`
- Estoque: `8083`
- Fiscal: `8084`
- Logística: `8085`
- Catálogo: `8086`