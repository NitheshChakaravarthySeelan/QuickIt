package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/db"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/kafka"
)

// StockHandler struct to manage stock operations
/* This is used to handle stock operations such as reserving and releasing stock.
It interacts with the database to update stock quantities and publishes events to Kafka for stock changes.
*/
type StockHandler struct {
	db *db.Postgres
	Producer *kafka.KafkaProducer
}

// NewStockHandler creates a new StockHandler with the provided database and Kafka producer
// Constructor function to initialize StockHandler with database and Kafka producer 
func NewStockHandler(db *db.Postgres, producer *kafka.KafkaProducer) *StockHandler {
	return &StockHandler{
		db: db,
		Producer: producer,
	}
}

// ReserveStock reserves stock for a given SKU and quantity
func (sh *StockHandler) ReserveStockHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.URL.Query().Get("sku")
	quantity := r.URL.Query().Get("quantity")

	if sku == "" || quantity == "" {
		log.Println("SKU and quantity are required")
		http.Error(w, "SKU and quantity are required", http.StatusBadRequest)
	}
	log.Printf("Reserving stock for SKU: %s, Quantity: %s", sku, quantity)

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Printf("Invalid quantity: %s", quantity)
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	reserveDB := sh.db.ReserveStock(ctx, sku, quantityInt)
	if reserveDB != nil {
		log.Printf("Failed to reserve stock: %v", reserveDB)
		http.Error(w, fmt.Sprintf("Failed to reserve stock: %v", reserveDB), http.StatusInternalServerError)
		return
	}
	log.Printf("Stock reserved successfully for SKU: %s, Quantity: %s", sku, quantity)
	
	// Publish stock reserved event to Kafka
	reserveEvent := kafka.StockReserveEvent{
		SKU: sku,
		ReservedQuantity: quantityInt,	
		ReservedAt: time.Now(),
		ReservedBy: "system", // This can be replaced with actual user information if available Need to implement user authentication
	}
	// Publish the event to Kafka
	log.Printf("Publishing stock reserved event for SKU: %s, Quantity: %d", reserveEvent.SKU, reserveEvent.ReservedQuantity)
	if err := sh.Producer.PublishStockReservedEvent(ctx, reserveEvent); err != nil {
		http.Error(w, fmt.Sprintf("Failed to publish stock reserved event: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("Stock reserved event published successfully for SKU: %s, Quantity: %d", reserveEvent.SKU, reserveEvent.ReservedQuantity)
	w.WriteHeader(http.StatusOK)
}

// ReleaseStock releases stock for a given SKU and quantity
func (sh *StockHandler) ReleaseStockHandler(w http.ResponseWriter, r *http.Request) {
	sku := r.URL.Query().Get("sku")
	quantity := r.URL.Query().Get("quantity")

	if sku == "" || quantity == "" {
		log.Println("SKU and quantity are required")
		http.Error(w, "SKU and quantity are required", http.StatusBadRequest)
		return
	}
	log.Printf("Releasing stock for SKU: %s, Quantity: %s", sku, quantity)

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Printf("Invalid quantity: %s", quantity)
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	releaseDB := sh.db.ReleaseStock(ctx, sku, quantityInt)
	if releaseDB != nil {
		log.Printf("Failed to release stock: %v", releaseDB)
		http.Error(w, fmt.Sprintf("Failed to release stock: %v", releaseDB), http.StatusInternalServerError)
		return
	}
	log.Printf("Stock released successfully for SKU: %s, Quantity: %s", sku, quantity)

	// Publish stock released event to Kafka
	releaseEvent := kafka.StockReleaseEvent{
		SKU: sku,
		ReleasedQuantity: quantityInt,
		ReleasedAt: time.Now(),
		ReleasedBy: "system", // This can be replaced with actual user information if available Need to implement user authentication
	}
	log.Printf("Publishing stock released event for SKU: %s, Quantity: %d", releaseEvent.SKU, releaseEvent.ReleasedQuantity)
	if err := sh.Producer.PublishStockReleasedEvent(ctx, releaseEvent); err != nil {
		http.Error(w, fmt.Sprintf("Failed to publish stock released event: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("Stock released event published successfully for SKU: %s, Quantity: %d", releaseEvent.SKU, releaseEvent.ReleasedQuantity)
	w.WriteHeader(http.StatusOK)
}