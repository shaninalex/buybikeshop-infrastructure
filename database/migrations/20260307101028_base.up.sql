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

CREATE OR REPLACE VIEW catalog.v_categories_tree AS
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
    category_id       bigint  NOT NULL,
    brand_id          bigint,
    description       varchar,
    short_description varchar,
    created_at        timestamp DEFAULT now(),
    updated_at        timestamp,

    FOREIGN KEY (brand_id) REFERENCES catalog.brands (id),
    FOREIGN KEY (category_id) REFERENCES catalog.categories (id)
);
CREATE INDEX products_title_idx ON catalog.products (title);


CREATE TABLE catalog.product_variants
(
    id          SERIAL PRIMARY KEY,
    product_id  bigint  NOT NULL,
    title       varchar NOT NULL,
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

-- INVENTORY

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
    id                 SERIAL PRIMARY KEY,
    sku                varchar UNIQUE,
    product_variant_id bigint NULL,

    FOREIGN KEY (product_variant_id)
        REFERENCES catalog.product_variants (id)
        ON DELETE SET NULL
);

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

-- PARTNERS


CREATE TABLE partners.roles
(
    id   SERIAL PRIMARY KEY,
    role varchar UNIQUE
);
-- 'Supplier', 'Customer', '3rd Party Worker'
CREATE TYPE partner_type AS ENUM ('person', 'company');

CREATE TABLE partners.partners
(
    id         SERIAL PRIMARY KEY,
    active     BOOLEAN      default True,
    type       partner_type default 'company',
    title      varchar not null,
    created_at timestamp    DEFAULT now()
);

CREATE TABLE partners.partner_roles
(
    role_id    bigint NOT NULL,
    partner_id bigint NOT NULL,
    FOREIGN KEY (partner_id) REFERENCES partners.partners (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES partners.roles (id) ON DELETE CASCADE
);

CREATE TABLE partners.partner_contacts
(
    id         SERIAL PRIMARY KEY,
    contacts   text   NOT NULL,
    partner_id bigint NOT NULL,
    created_at timestamp DEFAULT now(),

    FOREIGN KEY (partner_id) REFERENCES partners.partners (id) ON DELETE CASCADE
);

CREATE TABLE partners.suppliers
(
    id         SERIAL PRIMARY KEY,
    created_at timestamp DEFAULT now(),
    partner_id bigint NOT NULL,
    FOREIGN KEY (partner_id) REFERENCES partners.partners (id) ON DELETE CASCADE
);

