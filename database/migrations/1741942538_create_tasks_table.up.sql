CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_title ON tasks (title);
CREATE INDEX IF NOT EXISTS idx_status ON tasks (status);