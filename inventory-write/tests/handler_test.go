package db_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/db"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/handler"
	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/kafka"
)

func TestNewStockHandler(t *testing.T) {
	mockDB := &db.Postgres{}
	mockProducer := &kafka.KafkaProducer{}
	sh := handler.NewStockHandler(mockDB, mockProducer)
	if sh == nil {
		t.Error("Expected StockHandler to be initialized, got nil")
	}
}

func TestReserveStockHandlerBadRequest(t *testing.T) {
	mockDB := &db.Postgres{}
	mockProducer := &kafka.KafkaProducer{}
	sh := handler.NewStockHandler(mockDB, mockProducer)
	req := httptest.NewRequest("GET", "/reserve-stock", nil)
	w := httptest.NewRecorder()
	sh.ReserveStockHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
