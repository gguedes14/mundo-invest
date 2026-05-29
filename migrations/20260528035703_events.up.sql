CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id VARCHAR(255) UNIQUE NOT NULL,
    card_id VARCHAR(255) NOT NULL,
    cliente_email VARCHAR(255) NOT NULL,
    processed_at TIMESTAMP NOT NULL
);
