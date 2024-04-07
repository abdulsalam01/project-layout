CREATE TABLE IF NOT EXISTS channels(
    id SERIAL PRIMARY KEY,
    package_id BIGINT,
    name VARCHAR(255) NOT NULL,
    link VARCHAR(255) DEFAULT NULL,
    asset_url VARCHAR(255) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    created_by INT DEFAULT 0, -- Means by System.
    updated_by INT DEFAULT 0, -- Means by System.
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);