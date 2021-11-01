CREATE TABLE users (
    id INTEGER PRIMARY KEY, 
    email VARCHAR NOT NULL, 
    password VARCHAR NOT NULL
);

CREATE TABLE tasks (
    id INT PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    title VARCHAR NOT NULL
);