-- 1. Users кестесін жасау
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username TEXT NOT NULL UNIQUE,
                                     password TEXT NOT NULL,
                                     role TEXT DEFAULT 'user',
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP
);

-- 2. Todos кестесін жасау
CREATE TABLE IF NOT EXISTS todos (
                                     id SERIAL PRIMARY KEY,
                                     message TEXT NOT NULL,
                                     complete BOOLEAN DEFAULT FALSE,
                                     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP
    );
