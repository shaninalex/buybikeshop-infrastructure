CREATE TABLE inventory.warehouses
(
    id        SERIAL PRIMARY KEY,
    available bigint NOT NULL,
    address   text   NOT NULL
);

CREATE TABLE inventory.inventory
(
    id        SERIAL PRIMARY KEY,
    available bigint NOT NULL
);

ALTER TABLE catalog.product_variants
    ADD COLUMN inventory_id bigint;

ALTER TABLE catalog.product_variants
    ADD FOREIGN KEY (inventory_id)
        REFERENCES inventory.inventory (id)
        ON DELETE CASCADE;


CREATE TABLE inventory.deliveries
(
    id            SERIAL PRIMARY KEY,
    supplier_id   bigint NOT NULL,
    warehouse_id  bigint NOT NULL,
    delivery_date timestamp,
    status        varchar,
    created_at    timestamp DEFAULT now(),

    FOREIGN KEY (supplier_id)
        REFERENCES partners.suppliers (id),

    FOREIGN KEY (warehouse_id)
        REFERENCES inventory.warehouses (id)
);

CREATE TABLE inventory.delivery_items
(
    id          SERIAL PRIMARY KEY,
    delivery_id bigint  NOT NULL,
    variant_id  bigint  NOT NULL,
    quantity    integer NOT NULL,
    -- cost_price numeric,

    FOREIGN KEY (delivery_id) REFERENCES inventory.deliveries (id),
    FOREIGN KEY (variant_id) REFERENCES catalog.product_variants (id)
);
