
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS product (
    id_product SERIAL,
    name VARCHAR(16) NOT NULL,
    description VARCHAR(80),
    image_url VARCHAR(160),
    id_category INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id_product),
    FOREIGN KEY (id_category) REFERENCES category(id_category) ON DELETE SET NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS product;
