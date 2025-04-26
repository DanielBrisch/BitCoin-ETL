package models

import "time"

type CoinPrice struct {
	ID        uint      `gorm:"primaryKey"`
	Value     float64   `gorm:"column:value"`
	Cripto    string    `gorm:"column:cripto"`
	Coin      string    `gorm:"column:coin"`
	Timestamp time.Time `gorm:"column:timestamp"`
}
