CREATE TABLE IF NOT EXISTS approvals (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     owner_id int NOT NULL,
     url_id int NOT NULL,
     requester_id int NOT NULL,
     requester_reason TEXT,
     requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     is_approved int DEFAULT 0,
     approved_at TIMESTAMP,
     FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
     FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE ON UPDATE CASCADE,
     FOREIGN KEY (requester_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
)