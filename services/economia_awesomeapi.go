package services

import (
	"context"
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf(ctx.Err().Error() + " on request to API")
	default:
		return &cotacao.USDBRL, nil
	}
}
