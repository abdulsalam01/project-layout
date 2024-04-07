CREATE TABLE IF NOT EXISTS features(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    created_by INT DEFAULT 0, -- Means by System.
    updated_by INT DEFAULT 0, -- Means by System.
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);