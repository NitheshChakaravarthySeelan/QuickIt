package main

import (
	"fmt"
	"os"
	"context"
	"log"
	"net/http"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/db"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/kafka"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/handler"
)

var dbURL = os.Getenv("DATABASE_URL")
var kafkaBroker = os.Getenv("KAFKA_BROKER")
var kafkaTopic = os.Getenv("KAFKA_TOPIC")
var PORT = os.Getenv("PORT")


func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Service is running")
	})

	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKER is not set")
	}
	if kafkaTopic == "" {
		log.Fatal("KAFKA_TOPIC is not set")
	}	
	if PORT == "" {
		PORT = "8080"
	}


	ctx := context.Background()
	db, err := db.ConnectToDB(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	producer := kafka.NewKafkaProducer(kafkaBroker, kafkaTopic)
	defer producer.Close() // Ensure the producer is closed when done

	stockHandler := handler.NewStockHandler(&db, producer)
	if stockHandler == nil {
		log.Fatal("Failed to create StockHandler")
	}
	http.HandleFunc("/reserve-stock", stockHandler.ReserveStockHandler)
	http.HandleFunc("/release-stock", stockHandler.ReleaseStockHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}