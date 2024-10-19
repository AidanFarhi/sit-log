CREATE TABLE person (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
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
    FOREIGN KEY (adult_id) REFERENCES person(adult_id)
);