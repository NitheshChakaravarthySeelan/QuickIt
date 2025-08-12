from typing import List
from internal.models.stock import Stock
import logging
import aiohttp

warehouse_api_url = "https://api.example.com/warehouse"

async def get_stock_from_warehouse(warehouse_api_url: str = warehouse_api_url) -> List[Stock]:
    """Fetches stock data from the warehouse API asynchronously."""
    try:
        async with aiohttp.ClientSession() as session:
            async with session.get(f"{warehouse_api_url}/stock") as response:
                response.raise_for_status()  # Raises ClientResponseError for 4xx/5xx
                stock_data = await response.json()

                if not isinstance(stock_data, list):
                    logging.error("Invalid stock data format received: expected a list.")
                    return []

                return [Stock(**item) for item in stock_data]

    except aiohttp.ClientResponseError as e:
        logging.error(f"HTTP error fetching stock data: {e.status} {e.message} for {e.request_info.url}")
        return []
    except aiohttp.ClientError as e:
        logging.error(f"Request error fetching stock data: {e}")
        return []
    except Exception as e:
        # Catches other errors like JSON decoding or Stock model validation
        logging.error(f"An unexpected error occurred while processing stock data: {e}")
        return []
