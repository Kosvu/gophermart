CREATE TABLE users (
    login TEXT PRIMARY KEY,
    password TEXT NOT NULL,
    balance DECIMAL(12,2) NOT NULL DEFAULT 0 CHECK(balance >= 0)
);
