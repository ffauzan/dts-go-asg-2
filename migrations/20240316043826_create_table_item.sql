-- +goose Up
-- +goose StatementBegin
CREATE TABLE "items" (
    "id" SERIAL PRIMARY KEY,
    "order_id" INT NOT NULL,
    "code" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "quantity" INT NOT NULL,
    CONSTRAINT "fk_order_id" FOREIGN KEY ("order_id") REFERENCES "orders" ("id") ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "items";
-- +goose StatementEnd
