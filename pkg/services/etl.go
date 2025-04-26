package services

import (
	"context"
	"etlbitcoin/internal/database"
	"etlbitcoin/pkg/http"
	"etlbitcoin/pkg/models"
	"fmt"
	"strconv"
	"time"
)

func BitCoinDataTranslator(dados map[string]any) (*models.CoinPrice, error) {
	data, ok := dados["data"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected data format")
	}
	amount, err := strconv.ParseFloat(data["amount"].(string), 64)
	if err != nil {
		return nil, err
	}
	preco := &models.CoinPrice{
		Value:     amount,
		Cripto:    data["base"].(string),
		Coin:      data["currency"].(string),
		Timestamp: time.Now(),
	}
	return preco, nil
}

func PipelineBitcoin(ctx context.Context) {
	dadosJSON, err := http.GetBitcoinData(ctx)
	if err != nil {
		database.LogInsert("ERROR", fmt.Sprintf("Erro on extract data: %v", err))
		return
	}

	preco, err := BitCoinDataTranslator(dadosJSON)
	if err != nil {
		database.LogInsert("ERROR", fmt.Sprintf("Error when processing data: %v", err))
		return
	}

	database.SaveBitCoinPrice(preco)
	database.LogInsert("INFO", "Pipeline ETL BitCoin finish with sucess")
}
