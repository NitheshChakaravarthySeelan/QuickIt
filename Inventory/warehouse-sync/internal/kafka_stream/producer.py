import kafka
from kafka import KafkaProducer
from kafka.errors import KafkaError
import json
import logging
from typing import List
from internal.models.stock import Stock

producer = KafkaProducer(bootstrap_servers='kafka:9092')
topic = 'stock_updates'

def send_stock_updates(stock_data: List[Stock]):
    for item in stock_data:
        try:
            message = json.dumps(item.dict()).encode('utf-8')
            producer.send(topic, value=message)
            logging.info(f"Sent stock update for SKU: {item.sku}")
        except KafkaError as e:
            logging.error(f"Failed to send stock update for SKU: {item.sku} - {e}") 
        except Exception as e:
            logging.error(f"Unexpected error while sending stock update for SKU: {item.sku} - {e}")
    producer.flush()
    logging.info("All stock updates sent successfully.")

def close_producer():
    producer.close()
    logging.info("Kafka producer closed.")
