-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        price NUMERIC(10, 2),
        stock_count INT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
 