import pytest
import aiohttp
from aioresponses import aioresponses
from internal.erp.warehouse_api import get_stock_from_warehouse
from internal.models.stock import Stock

   
@pytest.mark.asyncio
async def test_get_stock_from_warehouse_success():
    url = "http://mock-warehouse-api.com/stock"

    mock_data = [
        {
            "sku": "SKU123",
            "quantity": 100,
            "location": "WH1"
        },
        {
            "sku": "SKU456",
            "quantity": 50,
            "location": "WH2"
        }
    ]

    with aioresponses() as m:
        m.get(url, payload=mock_data)

        result = await get_stock_from_warehouse(warehouse_api_url=url)

        assert isinstance(result, list)
        assert all(isinstance(stock, Stock) for stock in result)
        assert result[0].sku == "SKU123"
        assert result[1].quantity == 50


@pytest.mark.asyncio
async def test_get_stock_from_warehouse_empty():
    url = "http://mock-warehouse-api.com/stock"
    
    with aioresponses() as m:
        m.get(url, payload=[])

        result = await get_stock_from_warehouse(warehouse_api_url=url)
        assert result == []


@pytest.mark.asyncio
async def test_get_stock_from_warehouse_error():
    url = "http://mock-warehouse-api.com/stock"
    
    with aioresponses() as m:
        m.get(url, status=500)

        with pytest.raises(Exception):
            await get_stock_from_warehouse(warehouse_api_url=url)
