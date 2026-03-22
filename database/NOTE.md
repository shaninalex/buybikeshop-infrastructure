1) use shared schema:
```
shared
    countries
    currencies
    units
```

2) extend warehouse tables:
```sql
-- TODO:
-- for future development
-- supplier_sku       varchar,
-- supply_price       numeric(12,2),
-- currency           varchar(3),

-- How many days pass between placing a purchase order to supplier and receiving goods into warehouse.
-- lead_time_days     int,

-- Minimum quantity you must order from supplier in one purchase order.
-- min_order_qty      int,
```
