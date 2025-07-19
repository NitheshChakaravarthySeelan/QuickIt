package kafka

import "time"

type StockReserveEvent struct {
	SKU      string `json:"sku"`
	ReservedQuantity int    `json:"reserved_quantity"`
	ReservedAt time.Time `json:"reservedAtTime"`
	ReservedBy string `json:"reservedBy"`
}

type StockReleaseEvent struct {
	SKU      string `json:"sku"`
	ReleasedQuantity int    `json:"released_quantity"`
	ReleasedAt time.Time `json:"releasedAtTime"`
	ReleasedBy string `json:"releasedBy"`
}