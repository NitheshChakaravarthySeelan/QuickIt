use sqlx::{PgPool, Row};
use crate::utils::app_error::AppError;

pub struct InventoryRepo{
    pool: PgPool;
}

impl InventoryRepo {

    /// Creates a new instance of `InventoryRepo` with the provided database connection pool.
    pub fn new(pool: PgPool) -> Self {
        InventoryRepo { pool }
    }

    /// Ensures the tables exist - runss at startup
    pub async fn ensure_tables(&self) -> Result<(), AppError> {
        sqlx::query(
            "CREATE TABLE IF NOT EXISTS inventory (
                sku VARCHAR PRIMARY KEY,
                stock INTEGER NOT NULL DEFAULT 0
            )"
        )
        .execute(&self.pool)
        .await.map_err(|e| AppError::DatabaseError(e.to_string()))?;
        Ok(())
    }

    /// Fetches the stock for a given SKU returns quantity or 0 if not found
    pub async fn get_stock(&self, sku: &str) -> Result<i32, AppError> {
        let row: Option<(i32,)> = sqlx::query_as("SELECT stock FROM inventory WHERE sku = $1")
            .bind(sku)
            .fetch_optional(&self.pool)
            .await
            .map_err(|e| AppError::DatabaseError(e.to_string()))?;

        Ok(row.map_or(0, |(stock,)| stock))
    }
}
