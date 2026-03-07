CREATE TABLE inventory.warehouses
(
    id         SERIAL PRIMARY KEY,
    name       varchar   NOT NULL,
    address    text      NOT NULL,
    code       varchar UNIQUE,
    created_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE inventory.inventory_items
(
    id  SERIAL PRIMARY KEY,
    sku varchar UNIQUE
);

ALTER TABLE catalog.product_variants
    ADD COLUMN inventory_item_id bigint UNIQUE; -- should be not null

ALTER TABLE catalog.product_variants
    ADD FOREIGN KEY (inventory_item_id)
        REFERENCES inventory.inventory_items (id)
        ON DELETE CASCADE;

CREATE TABLE inventory.inventory_levels
(
    item_id      int    NOT NULL,
    warehouse_id int    NOT NULL,

    -- Physically in warehouse, ready to sell
    available    bigint NOT NULL DEFAULT 0,

    -- In transit / in delivery, not physically in warehouse yet
    incoming     bigint NOT NULL DEFAULT 0,

    -- Temporarily held (inspection, QA, allocated for a pick, or internal holds)
    reserved     bigint NOT NULL DEFAULT 0,

    -- Allocated to customer orders/carts but not shipped yet
    commited     bigint NOT NULL DEFAULT 0,

    PRIMARY KEY (item_id, warehouse_id),
    FOREIGN KEY (item_id) REFERENCES inventory.inventory_items (id),
    FOREIGN KEY (warehouse_id) REFERENCES inventory.warehouses (id)
);


--- DELIVERIES

CREATE TABLE inventory.deliveries
(
    id           SERIAL PRIMARY KEY,
    warehouse_id BIGINT    NOT NULL REFERENCES inventory.warehouses (id),
    supplier_id  BIGINT    NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE inventory.delivery_items
(
    delivery_id BIGINT REFERENCES inventory.deliveries (id),
    variant_id  BIGINT REFERENCES catalog.product_variants (id),
    quantity    BIGINT NOT NULL,

    PRIMARY KEY (delivery_id, variant_id)
);
