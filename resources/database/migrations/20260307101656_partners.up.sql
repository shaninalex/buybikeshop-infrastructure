CREATE TABLE partners.supplier_roles
(
    id   SERIAL PRIMARY KEY,
    role varchar UNIQUE
);

CREATE TABLE partners.suppliers
(
    id         SERIAL PRIMARY KEY,
    title      varchar UNIQUE NOT NULL,
    role_id    bigint         NOT NULL,
    created_at timestamp DEFAULT now(),

    FOREIGN KEY (role_id) REFERENCES partners.supplier_roles (id)
);

CREATE TABLE partners.suppliers_contacts
(
    id          SERIAL PRIMARY KEY,
    contacts    text   NOT NULL,
    supplier_id bigint NOT NULL,
    created_at  timestamp DEFAULT now(),

    FOREIGN KEY (supplier_id) REFERENCES partners.suppliers (id)
);
