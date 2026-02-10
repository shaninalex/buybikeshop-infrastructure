-- install extension for uuid primary keys
-- also make available to all database schemas.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" schema pg_catalog;

-- DEPRECATED: use "public" schema by default. Too mach schemas make another unnecessary level of complication
-- Inventory schema. Store all products, product variants, vendor procurement and vendors themselves
-- CREATE SCHEMA inventory;
-- CREATE SCHEMA market;

-- kratos is an identity provider. Will be used for customers.
CREATE SCHEMA customers;

-- kratos is an identity provider. Will be used company employees
CREATE SCHEMA employees;

-- databases for hydra
CREATE SCHEMA hydra;

-- databases for keto
CREATE SCHEMA keto;
