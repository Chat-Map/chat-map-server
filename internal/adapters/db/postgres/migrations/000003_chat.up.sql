CREATE TYPE "chat_t" AS enum ('private', 'group', 'channel');

CREATE TABLE "chat"
(
  id        BIGSERIAL PRIMARY KEY,
  chat_type chat_t NOT NULL
);

CREATE TABLE "chat_users"
(
  chat_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL
);

ALTER TABLE
  "chat_users"
  ADD
    FOREIGN KEY ("chat_id") REFERENCES "chat" ("id");

ALTER TABLE
  "chat_users"
  ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE
  "chat_users"
  ADD
    CONSTRAINT "chat_users_unique_id" UNIQUE (chat_id, user_id);

CREATE TABLE "messages"
(
  id         BIGSERIAL PRIMARY KEY,
  chat_id    BIGINT   NOT NULL,
  sender_id  BIGINT   NOT NULL,
  content    TEXT      NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
  "messages"
  ADD
    FOREIGN KEY ("chat_id") REFERENCES "chat" ("id");

ALTER TABLE
  "messages"
  ADD
    FOREIGN KEY ("sender_id") REFERENCES "users" ("id");
