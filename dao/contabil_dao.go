package dao

import (
	"context"
	"log"
	"time"

	"github.com/guirialli/dolar-api/database"
	"github.com/guirialli/dolar-api/models"
)

func ContabilAdd(ctx context.Context, contabil *models.Contabil) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-time.After(10 * time.Millisecond):
		database.DB.Save(contabil)
	}
}
