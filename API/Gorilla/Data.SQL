CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birthday VARCHAR(100) NOT NULL,
    phone_number VARCHAR(50) NOT NULL

);
CREATE TABLE addresses (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    district VARCHAR(50)NOT NULL,
    apartment VARCHAR(50),
    postal_code BIGINT NOT NULL,
    FOREIGN KEY(user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);