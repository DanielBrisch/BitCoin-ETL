package database

import (
	"etlbitcoin/pkg/models"
	"fmt"
	"os"
	"time"
)

func LogInsert(level, msg string) {
	db := DB
	entry := &models.LogEvent{
		Level:     level,
		Message:   msg,
		Timestamp: time.Now(),
	}
	if err := db.Create(entry).Error; err != nil {
		fmt.Fprintf(os.Stderr, "Error on save log: %v\n", err)
	}
}
