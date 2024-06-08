package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/guirialli/dolar-api/dao"
	"github.com/guirialli/dolar-api/models"
	"github.com/guirialli/dolar-api/services"
)

func ContabilGetAll(w http.ResponseWriter, r *http.Request) {
	cotacao, err := getCotacaoBRLContext()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeCotacaoDBContext(cotacao)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

	log.Println("Successfully send cotacao to client!")
}

func ContabilGetBid(w http.ResponseWriter, r *http.Request) {
	cotacao, err := getCotacaoBRLContext()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go writeCotacaoDBContext(cotacao)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao.Bid)

	log.Println("Successfully send cotacao to client!")
}

func getCotacaoBRLContext() (*models.Contabil, error) {
	ctxGetCotacao, cancelGetCotacao := context.WithTimeout(context.Background(), 200*time.Millisecond)
	cotacao, err := services.GetCotacaoDolarBRL(ctxGetCotacao)
	defer cancelGetCotacao()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return cotacao, nil
}

func writeCotacaoDBContext(cotacao *models.Contabil) error {
	ctxWriteContacao, cancelWriteContacao := context.WithTimeout(context.Background(), 10*time.Millisecond)
	err := dao.ContabilAdd(ctxWriteContacao, cotacao)
	defer cancelWriteContacao()
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
