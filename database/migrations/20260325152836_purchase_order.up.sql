CREATE TABLE inventory.purchase_orders
(
    id           SERIAL PRIMARY KEY,
    supplier_id  BIGINT    NOT NULL,
    warehouse_id BIGINT    NOT NULL,

    status       varchar   NOT NULL DEFAULT 'draft',
    -- draft | ordered | partially_received | received | cancelled

    expected_at  timestamp NULL,
    created_at   timestamp NOT NULL DEFAULT now(),
    updated_at   timestamp,

    FOREIGN KEY (supplier_id) REFERENCES partners.suppliers (id),
    FOREIGN KEY (warehouse_id) REFERENCES inventory.warehouses (id)
);

CREATE TABLE inventory.purchase_order_items
(
    id                 SERIAL PRIMARY KEY,
    purchase_order_id  BIGINT         NOT NULL,
    product_variant_id BIGINT         NOT NULL,

    quantity           BIGINT         NOT NULL,
    received_quantity  BIGINT         NOT NULL DEFAULT 0,

    price              NUMERIC(12, 2) NULL,

    FOREIGN KEY (purchase_order_id)
        REFERENCES inventory.purchase_orders (id)
        ON DELETE CASCADE,

    FOREIGN KEY (product_variant_id)
        REFERENCES catalog.product_variants (id)
);

--- DELIVERIES

CREATE TABLE inventory.deliveries
(
    id           SERIAL PRIMARY KEY,
    warehouse_id BIGINT    NOT NULL,
    supplier_id  BIGINT    NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT now(),

    FOREIGN KEY (warehouse_id) REFERENCES inventory.warehouses (id),
    FOREIGN KEY (supplier_id) REFERENCES partners.suppliers (id)
);

CREATE TABLE inventory.delivery_items
(
    delivery_id            BIGINT REFERENCES inventory.deliveries (id),
    purchase_order_item_id BIGINT NULL,
    quantity               BIGINT NOT NULL,

    PRIMARY KEY (delivery_id, purchase_order_item_id),
    FOREIGN KEY (purchase_order_item_id) REFERENCES inventory.purchase_order_items (id)
);
