-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
    id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    "state" varchar(255),
    "sum" DOUBLE PRECISION NOT NULL,
    "date" DATE NOT NULL

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
