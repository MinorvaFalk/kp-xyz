-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts(
    id serial PRIMARY KEY,
    customer_id bigint unsigned NOT NULL,
    created_at datetime DEFAULT now(),
    updated_at datetime,
    deleted_at datetime,
    CONSTRAINT fk_accounts_customers FOREIGN KEY (customer_id) REFERENCES customers (id)
        ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
