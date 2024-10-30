CREATE TABLE posts(
    id int UNIQUE AUTO_INCREMENT NOT NULL,
    title varchar(80) NOT NULL,
    content varchar(10000) NOT NULL
)