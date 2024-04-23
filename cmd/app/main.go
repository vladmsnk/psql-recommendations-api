package main

import (
	"log"
	"psqlRecommendationsApi/cmd/clients"
	"psqlRecommendationsApi/internal/config"
)

func main() {

	collectorClient, err := clients.NewCollectorClient(config.ConfigStruct.Collector)
	if err != nil {
		log.Fatal(err)
	}
	
	defer collectorClient.Close()
}
