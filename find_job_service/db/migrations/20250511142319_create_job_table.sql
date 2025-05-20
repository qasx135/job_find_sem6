-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    header TEXT NOT NULL,
    salary TEXT NOT NULL,
    experience TEXT NOT NULL,
    employment TEXT NOT NULL,
    schedule TEXT NOT NULL,
    work_format TEXT NOT NULL,
    working_hours TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS resume (
    id SERIAL PRIMARY KEY,
    about TEXT NOT NULL,
    experience TEXT NOT NULL,
    user_id INT NOT NULL UNIQUE ,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
