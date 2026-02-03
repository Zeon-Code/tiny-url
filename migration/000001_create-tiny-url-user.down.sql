REVOKE CONNECT ON DATABASE tiny_url FROM tiny_url;
REVOKE tiny_url_app FROM tiny_url;
DROP OWNED BY tiny_url_app;
DROP USER IF EXISTS tiny_url;
DROP ROLE IF EXISTS tiny_url_app;