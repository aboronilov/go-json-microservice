package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "BT")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(price)

	listenAddr := flag.String("listenaddr", ":3000", "the service is running")
	flag.Parse()
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJsonAPIServer(svc, *listenAddr)
	server.Run()
}
