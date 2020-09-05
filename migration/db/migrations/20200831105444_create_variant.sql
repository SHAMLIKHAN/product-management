
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS variant(
    id_variant SERIAL,
    name VARCHAR(16),
    max_retail_price FLOAT NOT NULL,
    discounted_price FLOAT,
    size VARCHAR(8),
    colour VARCHAR(8),
    id_product INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id_variant),
    FOREIGN KEY (id_product) REFERENCES product(id_product) ON DELETE CASCADE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS variant;
