ALTER TABLE warehouse.products
    ADD COLUMN vendor_id bigint NOT NULL;

ALTER TABLE warehouse.products
    ADD CONSTRAINT products_supplier_id_fk
    FOREIGN KEY (supplier_id) REFERENCES warehouse.supplier (id);

-- save links if posible

DROP TABLE warehouse.variant_suppliers;
