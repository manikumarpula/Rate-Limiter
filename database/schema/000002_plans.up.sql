-- Algorithm ENUMS
CREATE TYPE algorithm AS ENUM ('TOKEN_BUCKET', 'FIXED_WINDOW', 'LEAKY_BUCKET','SLIDING_WINDOW');


-- =========================
-- Table: plans
-- =========================
CREATE TABLE plans (
    id SERIAL PRIMARY KEY,
    account_id UUID NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,         
    rate_limit INTEGER NOT NULL,              
    burst_size INTEGER NOT NULL DEFAULT 0,    
    algorithm algorithm NOT NULL DEFAULT 'TOKEN_BUCKET', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_plans_name ON plans(name);
