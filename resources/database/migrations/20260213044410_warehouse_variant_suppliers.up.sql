ALTER TABLE warehouse.products
DROP COLUMN supplier_id;

CREATE TABLE warehouse.product_variant_suppliers
(
    product_variant_id bigint NOT NULL,
    supplier_id        bigint NOT NULL,

    PRIMARY KEY (product_variant_id, supplier_id),

    FOREIGN KEY (product_variant_id)
        REFERENCES warehouse.product_variants (id)
        ON DELETE CASCADE,

    FOREIGN KEY (supplier_id)
        REFERENCES warehouse.suppliers (id)
        ON DELETE CASCADE
);
