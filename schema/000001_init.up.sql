CREATE TABLE users 
(
    user_id SERIAL PRIMARY KEY,
    name varchar(256) not null,
    username varchar(256) not null unique,
    password_hash varchar(256) not null
);

CREATE TABLE todo_lists (
    list_id SERIAL PRIMARY KEY,
    title varchar(256) not null,
    description varchar(256)
);

CREATE TABLE users_lists (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(list_id) ON DELETE CASCADE
);

CREATE TABLE todo_items (
    item_id SERIAL PRIMARY KEY,
    title varchar(256) not null,
    description varchar(256),
    done boolean not null default false
);

CREATE TABLE list_items (
    id SERIAL PRIMARY KEY,
    item_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (item_id) REFERENCES todo_items(item_id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(list_id) ON DELETE CASCADE
);
