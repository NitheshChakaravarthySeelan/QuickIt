This service exists to keep the internal inventory database up to date with the real-world stock data from the external source like from ERP system, warehouse API's or partner vendors

Some update for the database update cause 
 - New goods arrive
 - Items are damaged or stolen 
 - Manual corrections are made

The inventory-write doesnt know automatically about these changes unless something pushes or pulls them. So the wareshouse-sync does the job of pulling the latest truth and updating the inventory.
