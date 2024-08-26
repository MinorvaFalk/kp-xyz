-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fees(
    id serial PRIMARY KEY,
    type integer NOT NULL,
    description varchar(255) NOT NULL,
    amount decimal(4,2) NOT NULL,
    created_at datetime DEFAULT now(),
    updated_at datetime,
    deleted_at datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fees;
-- +goose StatementEnd
