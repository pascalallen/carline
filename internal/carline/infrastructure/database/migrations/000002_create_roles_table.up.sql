CREATE TABLE roles
(
    id          CHAR(26) PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    modified_at TIMESTAMP
);

CREATE TABLE role_permissions
(
    role_id       CHAR(26) NOT NULL,
    permission_id CHAR(26) NOT NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);
