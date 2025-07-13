DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS popular_products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS anonymous_sessions;
DROP TABLE IF EXISTS request_metrics;
DROP TABLE IF EXISTS repeated_queries;

CREATE TABLE user_sessions
(
    session_id          UUID PRIMARY KEY,
    context             TEXT,
    last_active         TIMESTAMP DEFAULT NOW(),
    was_context_updated BOOLEAN   DEFAULT FALSE,
    user_id             UUID
);

CREATE TABLE anonymous_sessions
(
    session_id  UUID PRIMARY KEY,
    last_active TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE popular_products
(
    product_id  SERIAL PRIMARY KEY,
    name        VARCHAR(255),
    description TEXT,
    price       DECIMAL(10, 2),
    rating      DECIMAL(3, 2),
    category    VARCHAR(255),
    product_url VARCHAR(255),
    image_url   VARCHAR(255)
);

CREATE TABLE users
(
    user_id     UUID PRIMARY KEY,
    credentials TEXT,
    session_id  UUID,
    FOREIGN KEY (session_id) REFERENCES user_sessions (session_id) ON DELETE CASCADE
);

CREATE TABLE request_metrics
(
    id           SERIAL PRIMARY KEY,
    session_id   UUID NULL,
    request_type TEXT,
    success      BOOLEAN NOT NULL,
    duration     INT     NOT NULL
);

CREATE TABLE repeated_queries
(
    session_id    UUID PRIMARY KEY,
    similar_pairs INT NOT NULL,
    queries_count INT NOT NULL
);