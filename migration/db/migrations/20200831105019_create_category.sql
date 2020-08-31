
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS category (
    id_category SERIAL,
    name VARCHAR(16) NOT NULL,
    id_parent INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id_category),
    FOREIGN KEY (id_parent) REFERENCES category(id_category) ON DELETE SET NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS category;
