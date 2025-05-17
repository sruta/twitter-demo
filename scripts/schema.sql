CREATE TABLE `user`
(
    id       INT AUTO_INCREMENT PRIMARY KEY,
    email    VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    CONSTRAINT user_email_uindex UNIQUE (email),
    CONSTRAINT user_username_uindex UNIQUE (username),
    CONSTRAINT user_id_uindex UNIQUE (id)
);