-- CREATE TABLE orders (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     customer_name TEXT NOT NULL,
--     customer_phone TEXT NOT NULL,
--     total_amount INT NOT NULL,
--     currency TEXT NOT NULL DEFAULT 'XAF',
--     status TEXT NOT NULL DEFAULT 'PENDING',
--     created_at TIMESTAMP NOT NULL DEFAULT now()
-- );


CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_name TEXT NOT NULL,
    customer_phone TEXT NOT NULL,
    total_amount INT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'XAF',
    status TEXT NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
