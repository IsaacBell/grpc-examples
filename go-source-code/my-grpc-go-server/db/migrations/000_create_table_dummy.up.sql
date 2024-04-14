CREATE TABLE IF NOT EXISTS dummy (
  user_id       UUID PRIMARY KEY,
  user_name     TEXT NOT NULL,
  created_at    TIMESTAMPTZ,
  updated_at    TIMESTAMPTZ
);
