ALTER TABLE users
ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'developer',
ADD COLUMN department VARCHAR(50) NOT NULL DEFAULT 'TI';

-- Adicionar constraints
ALTER TABLE users
ADD CONSTRAINT valid_role CHECK (
    role IN ('admin', 'developer', 'devops', 'data-analyst', 'manager')
),
ADD CONSTRAINT valid_department CHECK (
    department IN ('data-analysis', 'TI', 'admin')
);