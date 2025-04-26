package database

import (
	"etlbitcoin/pkg/models"
	"fmt"
	"os"
	"time"
)

func CreateTables() {
	if err := DB.AutoMigrate(&models.CoinPrice{}, &models.LogEvent{}); err != nil {
		os.Exit(1)
		fmt.Println("erro ao migrar tabelas")
	}
	LogInsert("INFO", "tables created/verified successfully")
}

func SaveBitCoinPrice(price *models.CoinPrice) {
	if err := DB.Create(price).Error; err != nil {
		LogInsert("ERROR", fmt.Sprintf("Error on insert data: %v", err))
		return
	}
	LogInsert("INFO", fmt.Sprintf("[%s] BitCoin data saved", price.Timestamp.Format(time.RFC3339)))
}
