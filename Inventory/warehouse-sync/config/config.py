import os

# Warehouse API Configuration
WAREHOUSE_API_URL = os.getenv("WAREHOUSE_API_URL", "https://api.example.com/warehouse")

# Kafka Configuration
KAFKA_BOOTSTRAP_SERVERS = os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
KAFKA_STOCK_UPDATES_TOPIC = os.getenv("KAFKA_STOCK_UPDATES_TOPIC", "stock_updates")

# Database Configuration
DB_USER = os.getenv("DB_USER", "user")
DB_PASSWORD = os.getenv("DB_PASSWORD", "password")
DB_NAME = os.getenv("DB_NAME", "warehouse")
DB_HOST = os.getenv("DB_HOST", "localhost")
DB_PORT = int(os.getenv("DB_PORT", 5432))

# Logging Configuration
LOG_LEVEL = os.getenv("LOG_LEVEL", "INFO").upper()
