CREATE TABLE user_account
(
    id          uuid        DEFAULT uuid_generate_v4(),
    name        varchar(30) NOT NULL,
    email       varchar(80) NOT NULL,
    active      bool        DEFAULT FALSE,

    PRIMARY KEY (id),
    CONSTRAINT user_account_name_length CHECK (length(name) < 31),
    CONSTRAINT user_account_email_length CHECK (length(email) < 81)
);
