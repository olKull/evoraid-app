CREATE TABLE users
(
    id int primary key auto_increment not null,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id int primary key auto_increment not null,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id int primary key auto_increment not null,
    user_id int not null,
    list_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

CREATE TABLE todo_items
(
    id int primary key auto_increment not null,
    title varchar(255) not null,
    description varchar(255),
    done bit not null default false
);

CREATE TABLE lists_items
(
    id int auto_increment primary key,
    item_id int not null,
    list_id int not null,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) on DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) on DELETE CASCADE
);