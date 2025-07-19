package db_test

import (
	"testing"
	"time"

	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/kafka"
)

func TestStockReserveEventStruct(t *testing.T) {
	event := kafka.StockReserveEvent{
		SKU:              "SKU123",
		ReservedQuantity: 5,
		ReservedAt:       time.Now(),
		ReservedBy:       "tester",
	}
	if event.SKU != "SKU123" {
		t.Errorf("Expected SKU to be 'SKU123', got %s", event.SKU)
	}
	if event.ReservedQuantity != 5 {
		t.Errorf("Expected ReservedQuantity to be 5, got %d", event.ReservedQuantity)
	}
	if event.ReservedBy != "tester" {
		t.Errorf("Expected ReservedBy to be 'tester', got %s", event.ReservedBy)
	}
}

func TestStockReleaseEventStruct(t *testing.T) {
	event := kafka.StockReleaseEvent{
		SKU:              "SKU123",
		ReleasedQuantity: 3,
		ReleasedAt:       time.Now(),
		ReleasedBy:       "tester",
	}
	if event.SKU != "SKU123" {
		t.Errorf("Expected SKU to be 'SKU123', got %s", event.SKU)
	}
	if event.ReleasedQuantity != 3 {
		t.Errorf("Expected ReleasedQuantity to be 3, got %d", event.ReleasedQuantity)
	}
	if event.ReleasedBy != "tester" {
		t.Errorf("Expected ReleasedBy to be 'tester', got %s", event.ReleasedBy)
	}
}
