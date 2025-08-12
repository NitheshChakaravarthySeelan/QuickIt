from internal.models.stock import Stock
from typing import List
import logging
from internal.kafka_stream.producer import send_stock_updates, close_producer
import json
from internal.erp.warehouse_api import get_stock_from_warehouse
import asyncio
from internal.db.inventory_writer import update_stock_in_db, fetch_stock_from_db
import time
from internal.metrics.exporter import items_processed, last_sync_timestamp, sync_duration, sync_counter


async def sync_stock_updates():
    start_time = time.time()

    try:
        stock_data: List[Stock] = await get_stock_from_warehouse()
        
        if stock_data:
            processed_count = 0
            for item in stock_data:
                logging.info(f"Processing stock update for SKU: {item.sku}")
                try:
                    await update_stock_in_db(item)
                    processed_count += 1
                except Exception as e:
                    logging.error(f"Error updating stock in the database: {e}")

            # Update metrics
            if processed_count > 0:
                items_processed.set(processed_count)
                last_sync_timestamp.set(time.time())
                sync_counter.labels(status="success").inc()

                # Send to Kafka
                send_stock_updates(stock_data)
        else:
            logging.info("No stock updates to send.")
            sync_counter.labels(status="success").inc()
            items_processed.set(0)
            last_sync_timestamp.set(time.time())

    except Exception as e:
        logging.error(f"Error during stock synchronization: {e}")
        sync_counter.labels(status="failure").inc()

    finally:
        duration = time.time() - start_time
        sync_duration.observe(duration)
        close_producer()
        logging.info("Stock synchronization process completed.")
