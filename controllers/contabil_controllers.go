package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/guirialli/dolar-api/database"
	"github.com/guirialli/dolar-api/models"
)

func ContabilGet(w http.ResponseWriter, r *http.Request) {
	var cotacao models.USDBRL
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&cotacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Save(&cotacao.USDBRL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

	log.Println("Successfully send cotacao to client!")
}
