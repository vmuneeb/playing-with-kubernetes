package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type CurrencyResponse struct {
	Rate float64
	Host string
}

var currencyMap = map[string]float64{
	"usd": 1,
	"cad": 1.12,
	"inr": 72.3,
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	currency := r.FormValue("currency")
	value := currencyMap[currency]
	if value != 0 {
		resp := CurrencyResponse{
			Rate: value,
			Host: os.Getenv("HOSTNAME"),
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not found"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
