-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS info2 (
    id int8 NOT NULL GENERATED ALWAYS AS IDENTITY,
    state bool NOT NULL DEFAULT false,
    "date" date NOT NULL DEFAULT CURRENT_DATE,
    "name" varchar NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE info2;
-- +goose StatementEnd