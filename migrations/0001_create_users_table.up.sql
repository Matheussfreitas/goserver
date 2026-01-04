-- habilita extensão para gerar UUIDs
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    token TEXT,

    active BOOLEAN NOT NULL DEFAULT true,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- garante unicidade de email
CREATE UNIQUE INDEX idx_users_email ON users(email);
