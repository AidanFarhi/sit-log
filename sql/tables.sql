DROP TABLE IF EXISTS adult;
DROP TABLE IF EXISTS child;
DROP TABLE IF EXISTS adult_child_relation;
DROP TABLE IF EXISTS event;

CREATE TABLE adult (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE child (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    birthday TEXT NOT NULL
);

CREATE TABLE adult_child_relation (
    adult_id INTEGER,
    child_id INTEGER,
    relationship TEXT NOT NULL,
    PRIMARY KEY (adult_id, child_id),
    FOREIGN KEY (adult_id) REFERENCES adult(id),
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