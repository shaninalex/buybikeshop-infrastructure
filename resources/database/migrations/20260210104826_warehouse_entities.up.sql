CREATE TABLE warehouse.collections
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE
);

CREATE TABLE warehouse.brands
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE
);

CREATE TABLE warehouse.suppliers
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE
);

CREATE TABLE warehouse.categories
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(30) UNIQUE,
    parent_id BIGINT
);

CREATE OR REPLACE VIEW warehouse.v_categories_tree AS
(
SELECT c1.id,
       c1.title,
       parent.id    AS parent_id,
       parent.title AS parent_title
FROM warehouse.categories c1
        LEFT JOIN warehouse.categories parent ON c1.parent_id = parent.id
ORDER BY title
    );

CREATE INDEX idx_parent_id ON warehouse.categories (parent_id);

CREATE TABLE warehouse.products
(
    id                SERIAL PRIMARY KEY,
    title             varchar NOT NULL UNIQUE,
    collection_id     bigint,
    category_id       bigint  NOT NULL,
    brand_id          bigint,
    supplier_id       bigint,
    description       varchar,
    short_description varchar,
    created_at        timestamp DEFAULT now(),
    updated_at        timestamp,

    CONSTRAINT products_collection_id_fk FOREIGN KEY (collection_id) REFERENCES warehouse.collections (id),
    CONSTRAINT products_brand_id_fk FOREIGN KEY (brand_id) REFERENCES warehouse.brands (id),
    CONSTRAINT products_supplier_id_fk FOREIGN KEY (supplier_id) REFERENCES warehouse.suppliers (id),
    CONSTRAINT products_category_id_fk FOREIGN KEY (category_id) REFERENCES warehouse.categories (id)
);
CREATE INDEX products_title_idx ON warehouse.products (title);


CREATE TABLE warehouse.product_variants
(
    id          SERIAL          PRIMARY KEY,
    product_id  bigint          NOT NULL,
    title       varchar         NULL,
    description varchar         NULL,
    sku         varchar         NULL UNIQUE,
    barcode     varchar         NULL UNIQUE,
    created_at  timestamp       DEFAULT now(),
    updated_at  timestamp,

    CONSTRAINT store_product_variant_product_id_fk
        FOREIGN KEY (product_id)
            REFERENCES warehouse.products (id)
            ON DELETE CASCADE
);

CREATE INDEX store_products_variants_title_idx ON warehouse.product_variants (title);

CREATE TABLE warehouse.attributes
(
    id    SERIAL PRIMARY KEY,
    title varchar NOT NULL UNIQUE
);

CREATE TABLE warehouse.attributes_values
(
    id           SERIAL PRIMARY KEY,
    attribute_id bigint  NOT NULL,
    value        varchar NOT NULL,

    CONSTRAINT unique_attribute_value UNIQUE (attribute_id, value),
    CONSTRAINT store_products_attributes_values_attribute_id_fk
        FOREIGN KEY (attribute_id)
            REFERENCES warehouse.attributes (id)
            ON DELETE CASCADE
);

CREATE TABLE warehouse.product_variant_attributes
(
    id                 SERIAL PRIMARY KEY,
    product_variant_id bigint NOT NULL,
    attribute_id       bigint NOT NULL,
    value_id           bigint NOT NULL,

    CONSTRAINT unique_variant_attribute UNIQUE (product_variant_id, attribute_id),
    CONSTRAINT store_products_attributes_map_product_variant_id_fk
        FOREIGN KEY (product_variant_id)
            REFERENCES warehouse.product_variants (id)
            ON DELETE CASCADE,
    CONSTRAINT store_products_attributes_map_attribute_id_fk
        FOREIGN KEY (attribute_id)
            REFERENCES warehouse.attributes (id)
            ON DELETE CASCADE,
    CONSTRAINT store_products_attributes_map_value_id_fk
        FOREIGN KEY (value_id)
            REFERENCES warehouse.attributes_values (id)
            ON DELETE CASCADE
);

CREATE OR REPLACE VIEW warehouse.v_product_variant_attributes AS
(
SELECT map.product_variant_id AS variant_id,
       attr.id                AS attr_id,
       attr.title             AS attribute,
       val.value              AS value
FROM warehouse.product_variant_attributes map
         JOIN warehouse.attributes_values val ON map.value_id = val.id
         JOIN warehouse.attributes attr ON val.attribute_id = attr.id
    );

CREATE TABLE warehouse.product_variant_media
(
    id         SERIAL PRIMARY KEY,
    variant_id bigint  NOT NULL,
    media_type varchar NOT NULL,
    url        varchar NOT NULL,
    created_at timestamp DEFAULT NOW(),

    CONSTRAINT store_products_media_variant_id_fk
        FOREIGN KEY (variant_id)
            REFERENCES warehouse.product_variants (id)
            ON DELETE CASCADE,

    -- Maybe later will be convenient to create media_type table to hold additional types
    -- like: image, video, audio, 3d-presentation
    CONSTRAINT store_products_media_type CHECK (media_type IN ('image', 'video'))
);
