# ETL Bitcoin

Este projeto implementa um pipeline ETL (Extração, Transformação e Carga) que consulta a API pública da Coinbase a cada 15 segundos, trata os dados de preço do Bitcoin e persiste as informações em um banco de dados PostgreSQL.

---

## 🗂️ Estrutura de Diretórios

```
├── cmd
│   └── main.go            # Ponto de entrada da aplicação
│
├── internal/database      # Implementação de acesso a dados
│   ├── db.go              # Funções de conexão e migração do banco
│   ├── logger.go          # Lógica de gravação de logs no banco
│   └── repository.go      # Camada de persistência para entidades
│
├── pkg
│   ├── http
│   │   └── client.go      # Cliente HTTP para consumo da API Coinbase
│   ├── models
│   │   ├── coin_price.go  # Modelo para preço de criptomoeda
│   │   └── log_event.go   # Modelo para registro de logs
│   └── services
│       └── etl.go         # Lógica da pipeline ETL (extrair, transformar, carregar)
│
├── .env                   # Variáveis de ambiente (não comitar no repositório)
├── .gitignore             # Arquivos e pastas ignorados pelo Git
├── Dockerfile             # Imagem Docker para containerização
├── go.mod                 # Módulo Go e dependências
└── go.sum                 # Checksums das dependências
```

---

## 🚀 Tecnologias

- [Go](https://golang.org/) 1.20+
- [GORM](https://gorm.io/) (ORM para Go)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv) para carregamento de variáveis de ambiente
- Docker (opcional) para containerização

---

## ⚙️ Configuração

1. **Clone o repositório**
   ```bash
   git clone https://github.com/DanielBrisch/BitCoin-ETL.git
   cd BitCoin-ETL
   ```

2. **Arquivo `.env`**
   Crie um arquivo `.env` na raiz com as seguintes variáveis:
   ```dotenv
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=etl_bitcoin
   ```

3. **Inicialize o módulo e instale dependências**
   ```bash
   go mod download
   ```

---

## 🏃‍♂️ Execução

### Local

Para rodar localmente:

```bash
go run cmd/main.go
```

A cada 15 segundos, o programa:
1. Extrai o preço do Bitcoin na API da Coinbase.
2. Transforma o payload em estrutura `CoinPrice`.
3. Persiste o registro em `coin_prices` no PostgreSQL.
4. Registra eventos de log na tabela `logs`.

## 🔄 Estrutura de Banco de Dados

- **coin_prices**
  - `id` (serial, PK)
  - `value` (float)
  - `cripto` (varchar)
  - `coin` (varchar)
  - `timestamp` (timestamp)

- **logs**
  - `id` (serial, PK)
  - `level` (varchar)
  - `message` (text)
  - `timestamp` (timestamp)

---

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

