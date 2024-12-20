DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS child;
DROP TABLE IF EXISTS user_child_relation;
DROP TABLE IF EXISTS event;
DROP TABLE IF EXISTS session;

CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE child (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE user_child_relation (
    user_id INTEGER,
    child_id INTEGER,
    relationship TEXT NOT NULL,
    PRIMARY KEY (user_id, child_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (child_id) REFERENCES child(id)
);

CREATE TABLE event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    child_id INTEGER NOT NULL,
    timestamp TEXT DEFAULT (DATETIME('now')),
    type TEXT NOT NULL,
    description TEXT,
    start_time TEXT,
    end_time TEXT,
    duration TEXT,
    FOREIGN KEY (child_id) REFERENCES child(id)
);

CREATE TABLE session (
    token TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    created_at TEXT DEFAULT (DATETIME('now')),
    FOREIGN KEY (username) REFERENCES user(username)
);
