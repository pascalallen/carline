CREATE TABLE students
(
    id          CHAR(26) PRIMARY KEY,
    tag_number  VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    dismissed   BOOLEAN      NOT NULL,
    school_id   CHAR(26)     NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    modified_at TIMESTAMP,
    FOREIGN KEY (school_id) REFERENCES schools (id)
);
