# Bitcoin ETL

This project implements an ETL (Extract, Transform, Load) pipeline that polls the public Coinbase API every 15 seconds, processes Bitcoin price data, and persists the information to a PostgreSQL database.

---

## 🗂️ Directory Structure

```
├── cmd
│   └── main.go            # Application entry point
│
├── internal/database      # Data access implementation
│   ├── db.go              # Database connection and migration functions
│   ├── logger.go          # Logic for writing logs to the database
│   └── repository.go      # Persistence layer for entities
│
├── pkg
│   ├── http
│   │   └── client.go      # HTTP client for consuming the Coinbase API
│   ├── models
│   │   ├── coin_price.go  # Model for cryptocurrency price
│   │   └── log_event.go   # Model for log entries
│   └── services
│       └── etl.go         # ETL pipeline logic (extract, transform, load)
│
├── .env                   # Environment variables (do not commit)
├── .gitignore             # Files and folders ignored by Git
├── Dockerfile             # Docker image configuration
├── go.mod                 # Go module and dependencies
└── go.sum                 # Dependency checksums
```

---

## 🚀 Technologies

- [Go](https://golang.org/) 1.20+
- [GORM](https://gorm.io/) (ORM for Go)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv) for loading environment variables
- Docker (optional) for containerization

---

## ⚙️ Configuration

1. **Clone the repository**  
   ```bash
   git clone https://github.com/DanielBrisch/BitCoin-ETL.git
   cd BitCoin-ETL
   ```

2. **`.env` file**  
   Create a `.env` file in the project root with the following variables:  
   ```dotenv
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=etl_bitcoin
   ```

3. **Initialize the module and install dependencies**  
   ```bash
   go mod download
   ```

---

## 🏃‍♂️ Running

### Local

To run the application locally:

```bash
go run cmd/main.go
```

Every 15 seconds the program will:
1. Extract the current Bitcoin price from the Coinbase API.  
2. Transform the payload into the `CoinPrice` structure.  
3. Persist the record into the `coin_prices` table in PostgreSQL.  
4. Log events to the `logs` table.

---

## 🔄 Database Schema

- **coin_prices**
  - `id` (serial, PK)
  - `value` (float)
  - `crypto` (varchar)
  - `coin` (varchar)
  - `timestamp` (timestamp)

- **logs**
  - `id` (serial, PK)
  - `level` (varchar)
  - `message` (text)
  - `timestamp` (timestamp)

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

