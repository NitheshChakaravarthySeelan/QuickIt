use axum::{extract::{Path, State}, response::Json};
use serde_json::json;
use std::sync::Arc;
use crate::main::AppState;
use crate::utils::app_error::AppError;
use crate::handlers::health; // Added for test

/* This function is used to check the health of the service.
    It returns a JSON response with the status and message.
    It is typically used to verify that the service is up and running.
*/
pub async fn get_stock(
    Path(sku): Path<String>,
    State(state): State<Arc<AppState>>,
) -> Result<Json<serde_json::Value>, AppError> {
    let quantity = state.repo.get_stock(&sku).await?;

    Ok(Json(json!({
        "sku": sku,
        "quantity": quantity
    })))
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::db::InventoryRepo;
    use crate::kafka::KafkaPublisher;
    use sqlx::postgres::PgPoolOptions;
    use std::sync::Arc;

    #[tokio::test]
    async fn test_get_stock_returns_0_when_not_found() {
        // This test requires a running database instance via docker-compose.
        let pool = PgPoolOptions::new()
            .max_connections(1)
            .connect("postgres://inventory_user:inventory_pass@localhost:5432/inventory_db")
            .await
            .expect("Failed to connect to DB for test. Is `docker-compose up` running?");

        let app_state = Arc::new(AppState {
            repo: InventoryRepo::new(pool),
            kafka: KafkaPublisher::new("dummy:9092").unwrap(),
        });

        let response = get_stock(
            Path("non-existent-sku".to_string()),
            State(app_state),
        )
        .await
        .unwrap(); // The handler now returns a Result

        assert_eq!(response.0["quantity"], 0);
    }
}