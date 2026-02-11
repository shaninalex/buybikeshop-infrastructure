CREATE TABLE store_products_collection
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE
);

CREATE TABLE store_products_brand
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE
);

CREATE TABLE vendor
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE
);

CREATE TABLE store_products_categories
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(30) UNIQUE,
    level     INT NOT NULL DEFAULT 0,
    parent_id BIGINT
);

CREATE OR REPLACE VIEW v_products_categories AS
(
SELECT c1.id,
       c1.title,
       c1.level,
       parent.id    AS parent_id,
       parent.title AS parent_title
FROM store_products_categories c1
        LEFT JOIN store_products_categories parent ON c1.parent_id = parent.id
ORDER BY title
    );

CREATE INDEX idx_parent_id ON store_products_categories (parent_id);

CREATE TABLE store_products
(
    id                SERIAL PRIMARY KEY,
    title             varchar NOT NULL UNIQUE,
    collection_id     bigint,
    category_id       bigint  NOT NULL,
    brand_id          bigint,
    vendor_id         bigint, -- what if 2 vendors supply same product?
    description       varchar,
    short_description varchar,
    created_at        timestamp DEFAULT now(),
    updated_at        timestamp,

    CONSTRAINT products_collection_id_fk FOREIGN KEY (collection_id) REFERENCES store_products_collection (id),
    CONSTRAINT products_brand_id_fk FOREIGN KEY (brand_id) REFERENCES store_products_brand (id),
    CONSTRAINT products_vendor_id_fk FOREIGN KEY (vendor_id) REFERENCES vendor (id),
    CONSTRAINT products_category_id_fk FOREIGN KEY (category_id) REFERENCES store_products_categories (id)
);
CREATE INDEX products_title_idx ON store_products (title);


CREATE TABLE store_products_variants
(
    id          SERIAL          PRIMARY KEY,
    product_id  bigint          NOT NULL,
    title       varchar         NULL,
    description varchar         NULL,
    price       numeric(10, 2)  NOT NULL,
    sku         varchar         NULL UNIQUE,
    barcode     varchar         NULL UNIQUE,
    created_at  timestamp       DEFAULT now(),
    updated_at  timestamp,

    CONSTRAINT check_price_non_negative CHECK (price >= 0.00),
    CONSTRAINT store_product_variant_product_id_fk
        FOREIGN KEY (product_id)
            REFERENCES store_products (id)
            ON DELETE CASCADE
);

CREATE INDEX store_products_variants_title_idx ON store_products_variants (title);

CREATE TABLE store_products_attributes
(
    id    SERIAL PRIMARY KEY,
    title varchar NOT NULL UNIQUE
);

CREATE TABLE store_products_attributes_values
(
    id           SERIAL PRIMARY KEY,
    attribute_id bigint  NOT NULL,
    value        varchar NOT NULL,

    CONSTRAINT unique_attribute_value UNIQUE (attribute_id, value),
    CONSTRAINT store_products_attributes_values_attribute_id_fk
        FOREIGN KEY (attribute_id)
            REFERENCES store_products_attributes (id)
            ON DELETE CASCADE
);

CREATE TABLE store_products_attributes_map
(
    id                 SERIAL PRIMARY KEY,
    product_variant_id bigint NOT NULL,
    attribute_id       bigint NOT NULL,
    value_id           bigint NOT NULL,

    CONSTRAINT unique_variant_attribute UNIQUE (product_variant_id, attribute_id),
    CONSTRAINT store_products_attributes_map_product_variant_id_fk
        FOREIGN KEY (product_variant_id)
            REFERENCES store_products_variants (id)
            ON DELETE CASCADE,
    CONSTRAINT store_products_attributes_map_attribute_id_fk
        FOREIGN KEY (attribute_id)
            REFERENCES store_products_attributes (id)
            ON DELETE CASCADE,
    CONSTRAINT store_products_attributes_map_value_id_fk
        FOREIGN KEY (value_id)
            REFERENCES store_products_attributes_values (id)
            ON DELETE CASCADE
);

CREATE OR REPLACE VIEW v_products_attributes AS
(
SELECT map.product_variant_id AS variant_id,
       attr.id                AS attr_id,
       attr.title             AS attribute,
       val.value              AS value
FROM store_products_attributes_map map
         JOIN store_products_attributes_values val ON map.value_id = val.id
         JOIN store_products_attributes attr ON val.attribute_id = attr.id
    );

CREATE TABLE store_products_media
(
    id         SERIAL PRIMARY KEY,
    variant_id bigint  NOT NULL,
    media_type varchar NOT NULL,
    url        varchar NOT NULL,
    created_at timestamp DEFAULT NOW(),

    CONSTRAINT store_products_media_variant_id_fk
        FOREIGN KEY (variant_id)
            REFERENCES store_products_variants (id)
            ON DELETE CASCADE,

    -- Maybe later will be convenient to create media_type table to hold additional types
    -- like: image, video, audio, 3d-presentation
    CONSTRAINT store_products_media_type CHECK (media_type IN ('image', 'video'))
);