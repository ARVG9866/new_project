CREATE TABLE IF NOT EXISTS users  (
    id VARCHAR(27) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    preferred_locale VARCHAR(5) NOT NULL,

    PRIMARY KEY(id)
);


