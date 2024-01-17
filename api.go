package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/aboronilov/go-grpc-pricefetcher/types"
)

type APIfunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type JSONAPIServer struct {
	listenAddress string
	svc           PriceFetcher
}

func NewJsonAPIServer(svc PriceFetcher, address string) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddress: address,
		svc:           svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddress, nil)
}

func makeHTTPHandlerFunc(apiFn APIfunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")
	// if ticker == "" {
	// 	log.Error("No ticker query param is provided")
	// }

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	err = writeJSON(w, http.StatusOK, priceResponse)

	return err
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)

	return err
}
