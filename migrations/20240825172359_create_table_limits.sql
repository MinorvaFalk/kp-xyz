-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS limits (
    id serial PRIMARY KEY,
    account_id bigint unsigned NOT NULL,
    duration integer NOT NULL,
    initial_amount decimal(20,2) NOT NULL,
    current_amount decimal(20,2) NOT NULL,
    created_at datetime DEFAULT now(),
    updated_at datetime,
    deleted_at datetime,
    CONSTRAINT fk_limits_accounts FOREIGN KEY (account_id) REFERENCES accounts (id)
        ON UPDATE CASCADE    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS limits;
-- +goose StatementEnd
