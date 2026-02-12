-- install extension for uuid primary keys
-- also make available to all database schemas.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" schema pg_catalog;

-- kratos is an identity provider. Will be used for customers.
CREATE SCHEMA customers;

-- kratos is an identity provider. Will be used company employees
CREATE SCHEMA employees;

-- databases for hydra
CREATE SCHEMA hydra;

-- databases for keto ( permission management )
CREATE SCHEMA keto;

-- for warehouse
CREATE SCHEMA warehouse;
