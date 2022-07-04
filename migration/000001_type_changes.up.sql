CREATE TABLE users(
  id uuid NOT NULL PRIMARY KEY ,
  name varchar(50),
  number varchar(50)
);
CREATE TABLE posts(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    name text,
    title text,
    body text
    FOREIGN KEY(user_id)

);