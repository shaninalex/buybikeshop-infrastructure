CREATE TABLE IF NOT EXISTS applications
(
    id              uuid        DEFAULT uuid_generate_v4(),
    app_name        varchar(30) NOT NULL,
    active          bool        DEFAULT False,
    client_id       text        NOT NULL,
    client_secret   text        NOT NULL,
    url_home_page   text        NOT NULL,
    created_at      timestamp   DEFAULT now(),

    PRIMARY KEY (id),

    CONSTRAINT applications_name_length CHECK (length(app_name) < 31),
    CONSTRAINT applications_name_unique UNIQUE (app_name),
    CONSTRAINT applications_url_home_page_unique UNIQUE (url_home_page)
);


CREATE UNIQUE INDEX idx_applications_client_id_app_name ON applications(client_id, app_name);


CREATE TABLE IF NOT EXISTS application_redirect_uris
(
    id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    application_id  uuid NOT NULL,
    redirect_uri    text NOT NULL,
    active          boolean NOT NULL DEFAULT true,

    CONSTRAINT redirect_uri_unique UNIQUE (application_id, redirect_uri),

    FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE
);
