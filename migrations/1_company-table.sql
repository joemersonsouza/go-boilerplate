-- +goose Up

CREATE TABLE IF NOT EXISTS company
(
    id uuid NOT NULL,
    api_key uuid NOT NULL,
    CONSTRAINT company_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS company
    OWNER to admin;

-- +goose Down