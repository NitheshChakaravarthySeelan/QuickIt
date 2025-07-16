use axum::{response::{IntoResponse, Response}, http::StatusCode, Json};
use serde_json::json;

#[derive(Debug)]
pub enum AppError {
    DatabaseError(String),
    NotFoundError(String),
    KafkaError(String),
    BadRequestError(String),
    InternalServerError(String),
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let (status, error_message) = match self {
            AppError::DatabaseError(message) => (StatusCode::INTERNAL_SERVER_ERROR, format!("Database error: {}", message)),
            AppError::NotFoundError(message) => (StatusCode::NOT_FOUND, message),
            AppError::KafkaError(message) => (StatusCode::INTERNAL_SERVER_ERROR, format!("Kafka error: {}", message)),
            AppError::BadRequestError(message) => (StatusCode::BAD_REQUEST, message),
            AppError::InternalServerError(message) => (StatusCode::INTERNAL_SERVER_ERROR, message),
        };

        let body = Json(json!({
            "error": error_message,
        }));

        (status, body).into_response()
    }
}