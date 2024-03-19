-- +goose Up
-- +goose StatementBegin
CREATE TABLE "orders" (
    "id" SERIAL PRIMARY KEY,
    "customer_name" VARCHAR(255) NOT NULL,
    "ordered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "orders";
-- +goose StatementEnd
