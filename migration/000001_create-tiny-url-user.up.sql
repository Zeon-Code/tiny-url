CREATE ROLE tiny_url_app;
GRANT CONNECT ON DATABASE tiny_url TO tiny_url_app;
GRANT USAGE ON SCHEMA public TO tiny_url_app;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO tiny_url_app;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO tiny_url_app;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO tiny_url_app;

-- Password needs to be updated on production.
CREATE USER tiny_url WITH PASSWORD 'postgres';
GRANT tiny_url_app TO tiny_url;