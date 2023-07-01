CREATE TABLE
    IF
    NOT EXISTS category_posts (
     category_id INT,
     post_id INT,
    FOREIGN KEY ( category_id ) REFERENCES categories ( id ),
    FOREIGN KEY ( post_id ) REFERENCES posts ( id ),
    PRIMARY KEY ( category_id, post_id )
    );