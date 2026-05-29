CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE clientes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cliente_nome VARCHAR(255) NOT NULL,
    cliente_email VARCHAR(255) UNIQUE NOT NULL,
    tipo_solicitacao VARCHAR(255) NOT NULL,
    valor_patrimonio NUMERIC(15,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    prioridade VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
