CREATE TABLE IF NOT EXISTS "sessions" (
  id uuid PRIMARY KEY,
  user_id INTEGER NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
  "sessions"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");
