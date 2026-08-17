CREATE TABLE orders (
    number_order TEXT PRIMARY KEY,
    status TEXT NOT NULL,
    accrual DECIMAL(12,2),
    uploaded_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    user_login TEXT NOT NULL REFERENCES users(login)
);