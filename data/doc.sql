CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY, -- UUID as TEXT
    name TEXT NOT NULL,
    notesLocation TEXT NOT NULL,
	editor TEXT NOT NULL, 
	configPath TEXT NOT NULL

);

CREATE TABLE IF NOT EXISTS documents (
    id TEXT PRIMARY KEY,
    userID TEXT NOT NULL,
	directory TEXT,
    title TEXT,
	path TEXT,
    created_date TEXT, -- ISO 8601 format (used for date/time)
    keyword TEXT,
    FOREIGN KEY (userID) REFERENCES users(id) ON DELETE CASCADE
);

