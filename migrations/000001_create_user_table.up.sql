CREATE TABLE users(
  id serial primary key,
  username text not null,
  email text not null,
  password text not null,
  created_at timestamp default current_timestamp not null,
  updated_at timestamp default current_timestamp not null
);
