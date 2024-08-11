CREATE TABLE IF NOT EXISTS storePosition (
  id int8 PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  storeId varchar(500) NOT NULL,
  lat float8 NOT NULL,
  long float8 NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);