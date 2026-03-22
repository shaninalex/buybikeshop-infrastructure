ALTER TABLE catalog.product_variants
    ADD COLUMN price DECIMAL(12, 2) NOT NULL DEFAULT 0.00;

ALTER TABLE catalog.product_variants
    ADD COLUMN currency VARCHAR(3) NOT NULL DEFAULT 'eur';
