CREATE TABLE users (
    id serial primary key,
    username text not null,
    email text not null,
    password text not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

CREATE TABLE profiles (
    id serial primary key,
    user_id int not null references users(id) on delete cascade,
    bio text not null default '',
    avatar_url varchar(255) not null default '',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bio TEXT,
    avatar_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  type VARCHAR(50) NOT NULL,
  message TEXT NOT NULL,
  related_id INT NOT NULL, -- e.g., comments id
  is_read BOOLEAN DEFAULT FALSE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
);


-- HERE WE CREATE AND ADD INDEX
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_is_read ON notifications(is_read);
