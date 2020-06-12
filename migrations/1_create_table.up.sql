CREATE TABLE IF NOT EXISTS customers (
    id uuid PRIMARY KEY,
    phone VARCHAR(13) UNIQUE,
    email VARCHAR(30) UNIQUE
)

CREATE TABLE IF NOT EXISTS purchases (
    customer_id uuid NOT NULL REFERENCES customers(id),
    product_lists TEXT [], 
    purchases_total_price NUMERIC(5,5) NOT NULL,
    is_sent BOOLEAN NOT NULL DEFAULT FALSE
)

CREATE TABLE IF NOT EXISTS errors (
    error VARCHAR 
)