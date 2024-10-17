CREATE TABLE adult (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    name TEXT NOT NULL
);

CREATE TABLE child (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    parent_id INTEGER NOT NULL
    name TEXT NOT NULL
    FOREIGN KEY (parent_id) REFERENCES adult(id)
);

CREATE TABLE event (
    id INTEGER PRIMARY KEY AUTOINCREMENT
    adult_id INTEGER NOT NULL
    timestamp TEXT DEFAULT (datetime('now'))
    type TEXT NOT NULL
    description TEXT
    start_time TEXT
    end_time TEXT
    duration TEXT
    FOREIGN KEY (adult_id) REFERENCES adult(id)
);