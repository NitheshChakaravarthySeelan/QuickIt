import asyncpg
import logging
from typing import List
from internal.models.stock import Stock

async def connect_to_db():
    try:
        return await asyncpg.connect(user='user', password='password', database='warehouse', host='localhost', port=5432)
    except Exception as e:
        logging.error(f"Database connection error: {e}")
        return None
    
async def fetch_stock_from_db() -> List[Stock]:
    conn = await connect_to_db()
    
    if not conn:
        raise Exception("Failed to connect to the database")
    
    try:
        query = "SELECT sku, quantity, updated_at FROM stock"
        rows = await conn.fetch(query)
        
        if not rows:
            logging.info("No stock data found in the database.")
            return []
        
        return [Stock(sku=row['sku'], quantity=row['quantity'], updated_at=row['updated_at']) for row in rows]
    except asyncpg.PostgresError as e:
        logging.error(f"Database query error: {e}")
        return []
    finally:
        if conn:
            conn.close()

async def update_stock_in_db(stock_data: List[Stock]) -> List[Stock]:
    conn = await connect_to_db()
    
    if not conn:
        return []
    
    try:
        query = "INSERT INTO stock (sku, quantity, updated_at) " \
        "VALUES ($1, $2, $3) " \
        "ON CONFLICT (sku) " \
        "DO UPDATE SET quantity = EXCLUDED.quantity, " \
        "updated_at = EXCLUDED.updated_at"

        for item in stock_data:
            await conn.execute(query, item.sku, item.quantity, item.updated_at)
        logging.info("Stock data updated successfully in the database.")

    except asyncpg.PostgresError as e:
        logging.error(f"Database update error: {e}")
    finally:
        if conn:
            conn.close()

    return stock_data

