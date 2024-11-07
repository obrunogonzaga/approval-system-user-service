ALTER TABLE users
DROP CONSTRAINT valid_role,
DROP CONSTRAINT valid_department;

ALTER TABLE users
DROP COLUMN role,
DROP COLUMN department;