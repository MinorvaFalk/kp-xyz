-- +goose Up
-- +goose StatementBegin
insert into accounts (customer_id) values 
(1),
(2),
(3),
(4),
(5),
(6),
(7),
(8),
(9),
(10);
-- +goose StatementEnd

-- +goose NO TRANSACTION
-- +goose Down
SET FOREIGN_KEY_CHECKS = 0; 
TRUNCATE TABLE accounts;
SET FOREIGN_KEY_CHECKS = 1;

