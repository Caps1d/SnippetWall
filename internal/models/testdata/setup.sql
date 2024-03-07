-- Create snippets table
CREATE TABLE snippets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
    expires TIMESTAMPTZ NOT NULL
);

-- Create index on snippets.created
CREATE INDEX idx_snippets_created ON snippets(created);

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

-- Add unique constraint on users.email
ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

-- Insert data into users table
INSERT INTO users (name, email, hashed_password, created)
VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2022-01-01 10:00:00'::TIMESTAMPTZ
);

