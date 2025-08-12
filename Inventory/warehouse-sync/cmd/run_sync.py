import asyncio
import logging
from internal.sync.scheduler import scheduling

# Configure basic logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

async def main():
    """Main function to start the sync service."""
    logging.info("Starting the warehouse sync service...")
    await scheduling()

if __name__ == "__main__":
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        logging.info("Warehouse sync service stopped by user.")
