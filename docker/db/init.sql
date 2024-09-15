-- docker/db/init.sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    requests INT DEFAULT 0,
    last_reset TIMESTAMP
);
