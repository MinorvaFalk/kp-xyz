-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    account_id bigint unsigned NOT NULL,
    contract_number varchar(255) UNIQUE,
    asset_name varchar(255) NOT NULL,
    otr decimal(20, 2) NOT NULL,
    total_payment decimal(20, 2) NOT NULL,
    admin_fee decimal(20,2) NOT NULL,
    installment decimal(20, 2) NOT NULL,
    interest decimal(20, 2) NOT NULL,
    duration int NOT NULL,
    created_at datetime DEFAULT now(),
    updated_at datetime,
    deleted_at datetime,
    CONSTRAINT fk_transactions_accounts FOREIGN KEY (account_id) REFERENCES accounts (id)
        ON UPDATE CASCADE
) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
