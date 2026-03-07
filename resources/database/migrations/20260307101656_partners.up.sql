CREATE TABLE partners.partner_roles
(
    id   SERIAL PRIMARY KEY,
    role varchar UNIQUE
);

CREATE TABLE partners.partner
(
    id         SERIAL PRIMARY KEY,
    role_id    bigint  NOT NULL,
    title      varchar not null,
    created_at timestamp DEFAULT now(),

    FOREIGN KEY (role_id) REFERENCES partners.partner_roles (id)
);

CREATE TABLE partners.partner_contacts
(
    id         SERIAL PRIMARY KEY,
    contacts   text   NOT NULL,
    partner_id bigint NOT NULL,
    created_at timestamp DEFAULT now(),

    FOREIGN KEY (partner_id) REFERENCES partners.partner (id)
);

CREATE TABLE partners.suppliers
(
    id         SERIAL PRIMARY KEY,
    created_at timestamp DEFAULT now(),
    partner_id bigint NOT NULL,
    FOREIGN KEY (partner_id) REFERENCES partners.partner (id)
)
