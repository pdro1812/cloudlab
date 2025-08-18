-- currency-converter/init.sql
CREATE TABLE IF NOT EXISTS conversion_history (
    id SERIAL PRIMARY KEY,
    from_currency VARCHAR(3) NOT NULL,
    to_currency VARCHAR(3) NOT NULL,
    initial_amount NUMERIC(15, 4) NOT NULL,
    converted_amount NUMERIC(15, 4) NOT NULL,
    rate NUMERIC(15, 6) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);