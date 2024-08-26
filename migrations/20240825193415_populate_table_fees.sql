-- +goose Up
-- +goose StatementBegin
insert into fees (type, amount, description) values
(0, 0, 'admin fee'),
(1, 1.25, '1 month fee'),
(2, 1.75, '2 month fee'),
(3, 2.25, '3 month fee'),
(4, 2.5, '4 month fee');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE fees;
-- +goose StatementEnd
