CREATE TABLE IF NOT EXISTS accounts (
  id int8 PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  name varchar(255) NOT NULL
);