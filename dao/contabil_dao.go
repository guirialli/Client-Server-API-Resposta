package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/guirialli/dolar-api/database"
	"github.com/guirialli/dolar-api/models"
)

func ContabilAdd(ctx context.Context, contabil *models.Contabil) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf(ctx.Err().Error() + " on persist database")
	case <-time.After(10 * time.Millisecond):
		database.DB.Save(contabil)
		return nil
	}
}
