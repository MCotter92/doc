-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY, -- UUID as TEXT
    name TEXT NOT NULL,
    notes_location TEXT NOT NULL
);

-- Create documents table
CREATE TABLE IF NOT EXISTS documents (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    title TEXT,
    extension TEXT,
    location TEXT,
    created_date TEXT, -- ISO 8601 format (used for date/time)
    keyword TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

