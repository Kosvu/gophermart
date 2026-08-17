CREATE TABLE withdrawals (
    number_withdrawals TEXT PRIMARY KEY,
    amount DECIMAL(12,2) NOT NULL CHECK(amount > 0),
    user_login TEXT NOT NULL REFERENCES users(login),
    processed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);