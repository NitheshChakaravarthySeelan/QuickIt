use serde::{Deserialize, Serialize};

/// This is our domain data model (DTO)
/// It can serialize to JSON (for HTTP), deserialize from JSON or DB rows.
#[derive(Debug, Serialize, Deserialize)]
pub struct Stock {
    pub sku: String,
    pub quantity: i32,
}
