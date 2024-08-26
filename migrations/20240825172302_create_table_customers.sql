-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id serial PRIMARY KEY,
    nik varchar(16) UNIQUE,
    full_name varchar(255) NOT NULL,
    legal_name varchar(255) NOT NULL,
    birth_place varchar(255) NOT NULL,
    birth_date varchar(255) NOT NULL,
    salary decimal(20,2) NOT NULL,
    ktp_photo text NOT NULL,
    selfie_photo text NOT NULL,
    created_at datetime DEFAULT now(),
    updated_at datetime,
    deleted_at datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
