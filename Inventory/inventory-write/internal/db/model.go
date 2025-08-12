package db

type InventoryItem struct {
	SKU string `json:"sku"`
	AvailableQuantity int `json:"available_quantity"`
	ReservedQuantity int `json:"reserved_quantity"`
}