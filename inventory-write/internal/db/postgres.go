package db

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

/// PostgreSQL database connection and operations for inventory management struct to manage *pgxpool.Pool
type Postgres struct {
	pool *pgxpool.Pool	
}

// Connect establishes a connection to the PostgreSQL database using the provided connection string.

func ConnectToDB(ctx context.Context,dbUrl string) (Postgres, error){

	if dbUrl == "" {
		return Postgres{}, fmt.Errorf("DATABASE_URL is not set")
	}

	pool,err := pgxpool.Connect(ctx, dbUrl)

	if err != nil {
		return Postgres{}, fmt.Errorf("unable to connect to database: %v", err)
	}
	return Postgres{pool: pool}, nil
}

func (pg *Postgres) ReserveStock(ctx context.Context, sku string, quantity int) error {
	tx, err := pg.pool.Begin(ctx)

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx) // Ensure rollback if commit is not called

	const reserveQuery = `
	UPDATE inventory_items 
	SET reserved_quantity = reserved_quantity + $1 
	WHERE sku = $2 AND available_quantity >= $1`

	reservedStock, err := tx.Exec(ctx, reserveQuery, quantity, sku)
	if err != nil {
		return fmt.Errorf("failed to reserve stock: %v", err)
	}
	if reservedStock.RowsAffected() == 0 {
		return fmt.Errorf("not enough stock available for SKU %s", sku)
	}

	return tx.Commit(ctx)
}

func (pg *Postgres) ReleaseStock(ctx context.Context, sku string, quantity int) error {
	tx, err := pg.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx) // Ensure rollback if commit is not called

	const releaseQuery = `
	UPDATE inventory_items 
	SET reserved_quantity = reserved_quantity - $1 
	WHERE sku = $2`

	releasedStock, err := tx.Exec(ctx, releaseQuery, quantity, sku)
	if err != nil {
		return fmt.Errorf("failed to release stock: %v", err)
	}
	if releasedStock.RowsAffected() == 0 {
		return fmt.Errorf("no stock reserved for SKU %s", sku)
	}
	return tx.Commit(ctx)
}