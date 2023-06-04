CREATE TABLE "users"
(
  id         SERIAL PRIMARY KEY,
  first_name VARCHAR(50)  NOT NULL,
  last_name  VARCHAR(50)  NOT NULL,
  phone      VARCHAR(50)  NOT NULL,
  email      VARCHAR(100) NOT NULL,
  password   VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
  "users"
  ADD
    CONSTRAINT "users_unique_email" UNIQUE (email);

ALTER TABLE
  "users"
  ADD
    CONSTRAINT "users_unique_phone" UNIQUE (phone);

ALTER TABLE
  "users"
  ADD
    CONSTRAINT "phone_check" CHECK (phone ~ '^[0-9()+ -]+$');
