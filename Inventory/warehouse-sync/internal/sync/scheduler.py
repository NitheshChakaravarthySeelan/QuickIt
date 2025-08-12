import asyncio
import logging
import yaml
from apscheduler.schedulers.asyncio import AsyncIOScheduler
from internal.sync.job import sync_stock_updates

logging.basicConfig(level=logging.INFO)

def load_config():
    with open("config/scheduler.yaml", "r") as f:
        return yaml.safe_load(f)

async def safe_sync_stock_updates():
    try:
        await sync_stock_updates()
    except Exception as e:
        logging.error(f"Stock update failed: {e}")

async def scheduling():
    config = load_config()
    interval_minutes = config.get("interval_minutes", 10)

    scheduler = AsyncIOScheduler()
    scheduler.add_job(safe_sync_stock_updates, 'interval', minutes=interval_minutes)
    scheduler.start()

    logging.info(f"Scheduler started with interval {interval_minutes} minutes")
    
    try:
        while True:
            await asyncio.sleep(1)
    except KeyboardInterrupt:
        logging.info("Shutting down scheduler...")
        scheduler.shutdown()

