-- +goose Up
-- +goose StatementBegin
insert into customers (nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo) values 
('9191280503011873', 'Catarina Dziwisz', 'Catarina Dziwisz', 'China', '1990-07-31', 6219052, 'http://dummyimage.com/136x100.png/cc0000/ffffff', 'http://dummyimage.com/250x100.png/cc0000/ffffff'),
('7354424203937561', 'Kimmi Leyfield', 'Kimmi Leyfield', 'China', '1999-03-27', 6086609, 'http://dummyimage.com/217x100.png/ff4444/ffffff', 'http://dummyimage.com/151x100.png/dddddd/000000'),
('6552613005173647', 'Suzanne Cordes', 'Suzanne Cordes', 'China', '1998-06-14', 5698533, 'http://dummyimage.com/128x100.png/dddddd/000000', 'http://dummyimage.com/205x100.png/cc0000/ffffff'),
('1321114710171233', 'Corina O''Kuddyhy', 'Corina O''Kuddyhy', 'China', '1986-05-09', 9867267, 'http://dummyimage.com/185x100.png/cc0000/ffffff', 'http://dummyimage.com/113x100.png/5fa2dd/ffffff'),
('8198836212916070', 'Arvin Shevill', 'Arvin Shevill', 'Pakistan', '1987-11-30', 5444073, 'http://dummyimage.com/146x100.png/ff4444/ffffff', 'http://dummyimage.com/116x100.png/ff4444/ffffff'),
('6174712008451428', 'Alisander Rosenboim', 'Alisander Rosenboim', 'Venezuela', '1999-03-09', 9723144, 'http://dummyimage.com/165x100.png/ff4444/ffffff', 'http://dummyimage.com/125x100.png/dddddd/000000'),
('2189572612422643', 'Catherina Rubenovic', 'Catherina Rubenovic', 'Sweden', '1991-06-15', 8888766, 'http://dummyimage.com/243x100.png/cc0000/ffffff', 'http://dummyimage.com/136x100.png/cc0000/ffffff'),
('7297013105474160', 'Saba Trillo', 'Saba Trillo', 'China', '1995-06-22', 7935728, 'http://dummyimage.com/109x100.png/ff4444/ffffff', 'http://dummyimage.com/230x100.png/cc0000/ffffff'),
('2131710807877155', 'Aime Shinefield', 'Aime Shinefield', 'France', '1998-01-06', 6232609, 'http://dummyimage.com/225x100.png/ff4444/ffffff', 'http://dummyimage.com/175x100.png/cc0000/ffffff'),
('5327011511788673', 'Tomi Palfery', 'Tomi Palfery', 'Mexico', '1987-06-11', 8134944, 'http://dummyimage.com/131x100.png/5fa2dd/ffffff', 'http://dummyimage.com/160x100.png/ff4444/ffffff');
-- +goose StatementEnd

-- +goose NO TRANSACTION
-- +goose Down
SET FOREIGN_KEY_CHECKS = 0; 
TRUNCATE TABLE customers;
SET FOREIGN_KEY_CHECKS = 1;
