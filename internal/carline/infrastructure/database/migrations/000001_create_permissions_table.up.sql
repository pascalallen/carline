CREATE TABLE permissions
(
    id          CHAR(26) PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP    NOT NULL,
    modified_at TIMESTAMP
);
