package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJsonAPIServer(svc, *listenAddr)
	server.Run()

	// price, err := svc.FetchPrice(context.Background(), "BTC")
	// if err != nil {
	// 	log.Fatalf("Couldn't fetch price: %s", err)
	// }

	// fmt.Println(price)
}
