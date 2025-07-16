use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::ClientConfig;
use serde_json::Value;
use std::time::Duration;

/// Simple KafkaPublisher struct to encapsulate producer
pub struct KafkaPublisher {
    producer: FutureProducer,
}

impl KafkaPublisher {
    pub fn new(broker: &str) -> Self {
        let producer: FutureProducer = ClientConfig::new()
            .set("bootstrap.servers", broker)
            .create()
            .expect("Failed to create Kafka producer");

        KafkaPublisher { producer }
    }

    /// Sends a JSON payload to the given topic
    pub async fn send_event(&self, topic: &str, payload: &Value) {
        let json_str = payload.to_string();

        let delivery_status = self.producer.send(
            FutureRecord::to(topic)
                .payload(&json_str)
                .key("inventory"),
            Duration::from_secs(0),
        ).await;

        match delivery_status {
            Ok(delivery) => println!(" Kafka delivery: {:?}", delivery),
            Err((err, _)) => eprintln!(" Kafka error: {:?}", err),
        }
    }
}
