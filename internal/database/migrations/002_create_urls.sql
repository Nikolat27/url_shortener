CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    long_url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL,
    hashed_password TEXT,
    owner_id int NOT NULL,
    approval_required INTEGER DEFAULT 0,
    expiration_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
)