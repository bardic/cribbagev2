CREATE TABLE IF NOT EXISTS deck (
  id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  cards jsonb NOT NULL
);