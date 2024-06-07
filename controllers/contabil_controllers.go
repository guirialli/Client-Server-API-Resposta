package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/guirialli/dolar-api/dao"
	"github.com/guirialli/dolar-api/services"
)

func ContabilGet(w http.ResponseWriter, r *http.Request) {
	ctxGetCotacao, cancelGetCotacao := context.WithTimeout(context.Background(), 200*time.Millisecond)
	cotacao, err := services.GetCotacaoDolarBRL(ctxGetCotacao)
	defer cancelGetCotacao()

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctxWriteContacao, cancelWriteContacao := context.WithTimeout(context.Background(), 10*time.Millisecond)
	dao.ContabilAdd(ctxWriteContacao, cotacao)
	defer cancelWriteContacao()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

	log.Println("Successfully send cotacao to client!")
}
