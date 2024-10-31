CREATE TABLE users
(
    id            CHAR(26) PRIMARY KEY,
    first_name    VARCHAR(255) NOT NULL,
    last_name     VARCHAR(255) NOT NULL,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    modified_at   TIMESTAMP
);

CREATE TABLE user_schools
(
    user_id   CHAR(26) NOT NULL,
    school_id CHAR(26) NOT NULL,
    PRIMARY KEY (user_id, school_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (school_id) REFERENCES schools (id)
);

CREATE TABLE user_roles
(
    user_id CHAR(26) NOT NULL,
    role_id CHAR(26) NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
