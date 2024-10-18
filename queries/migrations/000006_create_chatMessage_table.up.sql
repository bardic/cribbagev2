CREATE TABLE IF NOT EXISTS chatMessage (
  id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  sender integer NOT NULL,
  message varchar(500) NOT NULL,
  timestamp timestamptz NOT NULL DEFAULT now()
);