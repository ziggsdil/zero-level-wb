CREATE TABLE IF NOT EXISTS orders
(
    order_uid VARCHAR(50) UNIQUE NOT NULL,
    data      JSONB              NOT NULL
)