-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    "id" INTEGER NOT NULL,
    "order" INTEGER NOT NULL,
    challenges INTEGER NOT NULL,
    "quantity" INTEGER NOT NULL


    created_at timestamp(0) NOT NULL DEFAULT now(),
    updated_at timestamp(0) NOT NULL DEFAULT now(),
    deleted_at timestamp(0)          DEFAULT NULL

    constraint fk_user_id foreign key (user_id) REFERENCES users (id) ON DELETE CASCADE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
