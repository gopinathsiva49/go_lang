CREATE TABLE IF NOT EXISTS `users` (
    id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    password varchar(255),
    first_name varchar(255),
    last_name varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX (email, first_name, last_name)
);