CREATE TABLE
    IF
    NOT EXISTS posts (
    id INT PRIMARY KEY,
    title VARCHAR ( 100 ) NOT NULL,
    TEXT VARCHAR ( 255 ) NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );