CREATE TABLE person (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE adult_child_relation (
    adult_id INTEGER NOT NULL,
    child_id INTEGER NOT NULL,
    relation TEXT NOT NULL,
    PRIMARY KEY (adult_id, child_id),
    FOREIGN KEY (adult_id) REFERENCES adult(id),
    FOREIGN KEY (child_id) REFERENCES child(id)
);

CREATE TABLE event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    adult_id INTEGER NOT NULL,
    child_id INTEGER NOT NULL,
    timestamp TEXT DEFAULT (datetime('now')),
    type TEXT NOT NULL,
    description TEXT,
    start_time TEXT,
    end_time TEXT,
    duration TEXT,
    FOREIGN KEY (adult_id, child_id) REFERENCES adult_child_relation(adult_id, child_id)
);