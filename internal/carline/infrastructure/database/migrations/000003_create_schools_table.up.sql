CREATE TABLE schools
(
    id          CHAR(26) PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    modified_at TIMESTAMP
);
