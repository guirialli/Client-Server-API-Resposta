package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/guirialli/dolar-api/models"
)

func GetCotacaoDolarBRL(ctx context.Context) (*models.Contabil, error) {
	var cotacao models.USDBRL
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&cotacao)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &cotacao.USDBRL, nil
	}
}
