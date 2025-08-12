# tests/test_sync_job.py
import pytest
import asyncio
from unittest.mock import AsyncMock, patch, MagicMock
from internal.sync.job import sync_stock_updates
from internal.models.stock import Stock

pytestmark = pytest.mark.asyncio

@patch("internal.sync.job.get_stock_from_warehouse", new_callable=AsyncMock)
@patch("internal.sync.job.update_stock_in_db", new_callable=AsyncMock)
@patch("internal.sync.job.send_stock_updates")
@patch("internal.sync.job.close_producer")
@patch("internal.sync.job.items_processed.set")
@patch("internal.sync.job.last_sync_timestamp.set")
@patch("internal.sync.job.sync_counter.labels")
@patch("internal.sync.job.sync_duration.observe")
async def test_successful_sync(
    mock_duration,
    mock_labels,
    mock_timestamp,
    mock_processed,
    mock_close,
    mock_send,
    mock_update,
    mock_get
):
    # Arrange
    mock_labels.return_value.inc = MagicMock()
    stock_list = [Stock(sku="ITEM123", quantity=10), Stock(sku="ITEM456", quantity=20)]
    mock_get.return_value = stock_list

    # Act
    await sync_stock_updates()

    # Assert
    assert mock_get.await_count == 1
    assert mock_update.await_count == len(stock_list)
    mock_send.assert_called_once_with(stock_list)
    mock_processed.assert_called_once_with(len(stock_list))
    mock_timestamp.assert_called()
    mock_labels.return_value.inc.assert_called_with()
    mock_close.assert_called_once()

@patch("internal.sync.job.get_stock_from_warehouse", new_callable=AsyncMock)
@patch("internal.sync.job.update_stock_in_db", new_callable=AsyncMock)
@patch("internal.sync.job.send_stock_updates")
@patch("internal.sync.job.close_producer")
@patch("internal.sync.job.items_processed.set")
@patch("internal.sync.job.last_sync_timestamp.set")
@patch("internal.sync.job.sync_counter.labels")
@patch("internal.sync.job.sync_duration.observe")
async def test_empty_erp_data(
    mock_duration,
    mock_labels,
    mock_timestamp,
    mock_processed,
    mock_close,
    mock_send,
    mock_update,
    mock_get
):
    mock_labels.return_value.inc = MagicMock()
    mock_get.return_value = []

    await sync_stock_updates()

    mock_send.assert_not_called()
    mock_update.assert_not_called()
    mock_processed.assert_called_once_with(0)
    mock_timestamp.assert_called()
    mock_labels.return_value.inc.assert_called_with()

@patch("internal.sync.job.get_stock_from_warehouse", new_callable=AsyncMock)
@patch("internal.sync.job.update_stock_in_db", new_callable=AsyncMock)
@patch("internal.sync.job.send_stock_updates")
@patch("internal.sync.job.close_producer")
@patch("internal.sync.job.items_processed.set")
@patch("internal.sync.job.last_sync_timestamp.set")
@patch("internal.sync.job.sync_counter.labels")
@patch("internal.sync.job.sync_duration.observe")
async def test_partial_failure_in_db_update(
    mock_duration,
    mock_labels,
    mock_timestamp,
    mock_processed,
    mock_close,
    mock_send,
    mock_update,
    mock_get
):
    mock_labels.return_value.inc = MagicMock()
    stock_list = [Stock(sku="ITEM123", quantity=10), Stock(sku="ITEM456", quantity=20)]
    mock_get.return_value = stock_list

    async def side_effect(stock):
        if stock.sku == "ITEM456":
            raise Exception("DB Error")
        return

    mock_update.side_effect = side_effect

    await sync_stock_updates()

    assert mock_update.await_count == 2
    mock_send.assert_called_once_with([Stock(sku="ITEM123", quantity=10), Stock(sku="ITEM456", quantity=20)])

@patch("internal.sync.job.get_stock_from_warehouse", new_callable=AsyncMock)
@patch("internal.sync.job.sync_counter.labels")
@patch("internal.sync.job.sync_duration.observe")
@patch("internal.sync.job.close_producer")
async def test_erp_failure(
    mock_close,
    mock_duration,
    mock_labels,
    mock_get
):
    mock_labels.return_value.inc = MagicMock()
    mock_get.side_effect = Exception("ERP Down")

    await sync_stock_updates()

    mock_labels.return_value.inc.assert_called_with()
