-- Status ENUM
CREATE TYPE status AS ENUM ('ACTIVE', 'INACTIVE');
-- Product ENUM
CREATE TYPE product AS ENUM ('PRODUCT_1', 'PRODUCT_2', 'PRODUCT_3');
-- =========================
-- Table: api_keys
-- =========================
CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    key UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    status status NOT NULL DEFAULT 'ACTIVE',
    created_by VARCHAR(100),
    last_used TIMESTAMP,
    requests_consumed BIGINT DEFAULT 0,
    product product NOT NULL DEFAULT 'PRODUCT_1',
    plan_id INT NOT NULL REFERENCES plans(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);