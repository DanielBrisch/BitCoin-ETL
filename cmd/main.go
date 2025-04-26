package main

import (
	"context"
	"etlbitcoin/internal/database"
	"etlbitcoin/pkg/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("/.env"); err != nil {
		fmt.Println("Erro ao carregar .env")
	}

	database.ConnectDB()
	database.CreateTables()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			database.LogInsert("INFO", "Process interrupted by user. Ending")
			return
		case <-ticker.C:
			services.PipelineBitcoin(ctx)
		}
	}
}
