CREATE TABLE IF NOT EXISTS cards (
  id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  value integer NOT NULL,
  suit integer NOT NULL,
  art varchar(500) NOT NULL
);