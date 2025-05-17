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

CREATE TABLE follower
(
    follower_id INT      NOT NULL,
    followed_id INT      NOT NULL,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followed_id),
    CONSTRAINT follower_follower_id_fk FOREIGN KEY (follower_id) REFERENCES user (id) ON DELETE CASCADE,
    CONSTRAINT follower_followed_id_fk FOREIGN KEY (followed_id) REFERENCES user (id) ON DELETE CASCADE
)