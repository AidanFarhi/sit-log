CREATE TABLE adult (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE child (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parent_id INTEGER,
    name TEXT NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES adult(id)
);

CREATE TABLE event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    adult_id INTEGER NOT NULL,
    child_id INTEGER NOT NULL,
    timestamp TEXT DEFAULT (DATETIME('now')),
    type TEXT NOT NULL,
    description TEXT,
    start_time TEXT,
    end_time TEXT,
    duration TEXT,
    FOREIGN KEY (adult_id) REFERENCES adult(id),
    FOREIGN KEY (child_id) REFERENCES child(id)
);