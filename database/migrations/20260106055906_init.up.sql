CREATE TABLE IF NOT EXISTS identities
(
    id              uuid        DEFAULT uuid_generate_v4(),
    fullname        varchar(30) NOT NULL,
    email           varchar(80) NOT NULL,
    active          bool        DEFAULT False,
    created_at      timestamp   DEFAULT now(),

    PRIMARY KEY (id),

    CONSTRAINT identity_fullname_length CHECK (length(fullname) < 31),
    CONSTRAINT identity_email_unique UNIQUE (email),
    CONSTRAINT identity_email_length CHECK (length(email) < 81)
);
