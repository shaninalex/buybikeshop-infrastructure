CREATE TABLE catalog.collections
(
    id   SERIAL PRIMARY KEY,
    name varchar(30) UNIQUE
);

CREATE TABLE catalog.brands
(
    id    SERIAL PRIMARY KEY,
    title varchar UNIQUE
);

CREATE TABLE catalog.categories
(
    id        SERIAL PRIMARY KEY,
    title     varchar(30) UNIQUE,
    parent_id bigint
);

CREATE
OR REPLACE VIEW catalog.v_categories_tree AS
(
    SELECT c1.id,
        c1.title,
        parent.id    AS parent_id,
        parent.title AS parent_title
    FROM catalog.categories c1
        LEFT JOIN catalog.categories parent ON c1.parent_id = parent.id
    ORDER BY title
);

CREATE INDEX idx_parent_id ON catalog.categories (parent_id);

CREATE TABLE catalog.products
(
    id                SERIAL PRIMARY KEY,
    title             varchar NOT NULL UNIQUE,
    collection_id     bigint,
    category_id       bigint  NOT NULL,
    brand_id          bigint,
    description       varchar,
    short_description varchar,
    created_at        timestamp DEFAULT now(),
    updated_at        timestamp,

    FOREIGN KEY (collection_id) REFERENCES catalog.collections (id),
    FOREIGN KEY (brand_id) REFERENCES catalog.brands (id),
    FOREIGN KEY (category_id) REFERENCES catalog.categories (id)
);
CREATE INDEX products_title_idx ON catalog.products (title);

CREATE TABLE catalog.product_variants
(
    id          SERIAL PRIMARY KEY,
    product_id  bigint NOT NULL,
    title       varchar NULL,
    description varchar NULL,
    sku         varchar NULL UNIQUE,
    barcode     varchar NULL UNIQUE,
    created_at  timestamp DEFAULT now(),
    updated_at  timestamp,

    FOREIGN KEY (product_id)
        REFERENCES catalog.products (id)
        ON DELETE CASCADE
);

CREATE INDEX catalog_products_variants_title_idx ON catalog.product_variants (title);

CREATE TABLE catalog.attributes
(
    id    SERIAL PRIMARY KEY,
    title varchar NOT NULL UNIQUE
);

CREATE TABLE catalog.attributes_values
(
    id           SERIAL PRIMARY KEY,
    attribute_id bigint  NOT NULL,
    value        varchar NOT NULL,

    CONSTRAINT unique_attribute_value UNIQUE (attribute_id, value),
    FOREIGN KEY (attribute_id)
        REFERENCES catalog.attributes (id)
        ON DELETE CASCADE
);

CREATE TABLE catalog.product_variant_attributes
(
    id                 SERIAL PRIMARY KEY,
    product_variant_id bigint NOT NULL,
    attribute_id       bigint NOT NULL,
    value_id           bigint NOT NULL,

    CONSTRAINT unique_variant_attribute UNIQUE (product_variant_id, attribute_id),
    FOREIGN KEY (product_variant_id)
        REFERENCES catalog.product_variants (id)
        ON DELETE CASCADE,
    FOREIGN KEY (attribute_id)
        REFERENCES catalog.attributes (id)
        ON DELETE CASCADE,
    FOREIGN KEY (value_id)
        REFERENCES catalog.attributes_values (id)
        ON DELETE CASCADE
);

CREATE
OR REPLACE VIEW catalog.v_product_variant_attributes AS
(
    SELECT map.product_variant_id AS variant_id,
       attr.id                AS attr_id,
       attr.title             AS attribute,
       val.value              AS value
    FROM catalog.product_variant_attributes map
         JOIN catalog.attributes_values val ON map.value_id = val.id
         JOIN catalog.attributes attr ON val.attribute_id = attr.id
);

CREATE TABLE catalog.product_variant_media
(
    id         SERIAL PRIMARY KEY,
    variant_id bigint  NOT NULL,
    media_type varchar NOT NULL,
    url        varchar NOT NULL,
    created_at timestamp DEFAULT NOW(),

    FOREIGN KEY (variant_id)
        REFERENCES catalog.product_variants (id)
        ON DELETE CASCADE,

    -- Maybe later will be convenient to create media_type table to hold additional types
    -- like: image, video, audio, 3d-presentation
    CONSTRAINT catalog_products_media_type CHECK (media_type IN ('image', 'video'))
);
