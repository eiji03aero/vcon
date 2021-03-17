
-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL
);

-- +migrate Down
DROP TABLE users;
