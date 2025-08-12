from prometheus_client import Counter, Gauge, Summary

# Counter with 'status' label
sync_counter = Counter(
    'warehouse_sync_total',
    'Total number of warehouse sync operations',
    ['status']  # 'success' or 'failure'
)

# Summary for measuring duration
sync_duration = Summary(
    'warehouse_sync_duration_seconds',
    'Duration of warehouse sync operations in seconds'
)

# Gauge for number of items processed
items_processed = Gauge(
    'warehouse_sync_items_processed',
    'Number of items processed in the last sync'
)

# Gauge for the last successful sync timestamp (Unix time)
last_sync_timestamp = Gauge(
    'warehouse_sync_last_sync_timestamp',
    'Unix timestamp of the last successful sync'
)
