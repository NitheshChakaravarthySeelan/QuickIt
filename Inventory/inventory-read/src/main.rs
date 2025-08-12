use axum::{Router, routing::get};
use dotenv::dotenv;
use sqlx::PgPool;
use std::{env, net::SocketAddr, sync::Arc};

mod db;
mod handlers;
mod kafka;
mod models;
mod utils;

use db::InventoryRepo;
use kafka::KafkaPublisher;
use handlers::{health, get_stock};

/// Shared state struct for Dependency Injection
pub struct AppState {
    pub repo: InventoryRepo,
    pub kafka: KafkaPublisher,
}

#[tokio::main]
async fn main() {
    dotenv().ok();
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let kafka_broker = env::var("KAFKA_BROKER").unwrap_or_else(|_| "localhost:9092".into());

    // Set up Postgres
    let pool = PgPool::connect(&database_url)
        .await
        .expect("Failed to connect to Postgres");

    let repo = InventoryRepo::new(pool);
    repo.ensure_tables().await;

    // Set up Kafka
    let kafka = KafkaPublisher::new(&kafka_broker);

    // Shared state (DI)
    let app_state = Arc::new(AppState { repo, kafka });

    // Build Axum app with routes and state
    let app = Router::new()
        .route("/health", get(health))
        .route("/stock/:sku", get(get_stock))
        .with_state(app_state);

    let addr = SocketAddr::from(([0, 0, 0, 0], 3000));
    println!("Inventory-Read running on {}", addr);

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}