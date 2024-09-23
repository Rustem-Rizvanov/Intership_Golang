CREATE TABLE IF NOT EXISTS request_statistics (
    id SERIAL PRIMARY KEY,
    processing_time FLOAT NOT NULL,
    outcome VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
