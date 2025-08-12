from pydantic import BaseModel
from datetime import datetime

class Stock(BaseModel):
    sku: str
    quantity: int 
    updated_at: datetime