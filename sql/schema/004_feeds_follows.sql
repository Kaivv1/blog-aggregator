-- +goose Up 
CREATE TABLE feeds_follows (
  id UUID PRIMARY KEY, 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP NOT NULL,  
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE(user_id, feed_id)
);

CREATE INDEX feeds_follows_user_idx ON feeds_follows(user_id);
CREATE INDEX feeds_follows_feed_idx ON feeds_follows(feed_id);

-- +goose Down
DROP INDEX feeds_follows_user_idx;
DROP INDEX feeds_follows_feed_idx;

DROP TABLE feeds_follows;
