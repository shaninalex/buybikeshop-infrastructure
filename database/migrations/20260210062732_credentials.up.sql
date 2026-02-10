-- NOTES:
--  - provider should be as a separate table and instead provider text field - provider_id foreign key

CREATE TABLE IF NOT EXISTS credentials
(
    id                  uuid            DEFAULT uuid_generate_v4(),
    identity_id         uuid            NOT NULL,
    provider            varchar(20)     NOT NULL,
    provider_user_id    varchar(36)     NULL,
    email               varchar(80)     NULL, -- make reasonable length of email text field
    pwd_hash            text            NULL,
    created_at          timestamp       DEFAULT now(),

    PRIMARY KEY (id),

    CONSTRAINT credentials_identity_id_fk FOREIGN KEY (identity_id) REFERENCES identities (id)
);
