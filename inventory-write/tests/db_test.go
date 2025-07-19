package db_test

import (
	"context"
	"testing"

	"github.com/NitheshChakaravarthySeelan/QuickIt/inventory-write/internal/db"
)

func TestInventoryItemStruct(t *testing.T) {
	item := db.InventoryItem{
		SKU:               "SKU123",
		AvailableQuantity: 100,
		ReservedQuantity:  10,
	}
	if item.SKU != "SKU123" {
		t.Errorf("Expected SKU to be 'SKU123', got %s", item.SKU)
	}
	if item.AvailableQuantity != 100 {
		t.Errorf("Expected AvailableQuantity to be 100, got %d", item.AvailableQuantity)
	}
	if item.ReservedQuantity != 10 {
		t.Errorf("Expected ReservedQuantity to be 10, got %d", item.ReservedQuantity)
	}
}

func TestConnectToDBEmptyURL(t *testing.T) {
	ctx := context.Background()
	_, err := db.ConnectToDB(ctx, "")
	if err == nil {
		t.Error("Expected error when DATABASE_URL is empty, got nil")
	}
}
