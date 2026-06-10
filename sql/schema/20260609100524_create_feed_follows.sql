-- +goose Up
CREATE TABLE feed_follows (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  feed_id UUID REFERENCES feeds ON DELETE CASCADE NOT NULL,
  user_id UUID REFERENCES users ON DELETE CASCADE NOT NULL
);
CREATE UNIQUE INDEX feed_followers_unique_idx ON feed_follows(user_id, feed_id);

-- +goose Down
DROP TABLE feed_follows;
