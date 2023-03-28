CREATE TABLE IF NOT EXISTS `sessions` (
    id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT,
    uuid longtext,
    password varchar(255),
    user_id bigint,
    salt varchar(191),
    session_key varchar(191),
    auth_token TEXT,
    expires_at DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX (user_id, salt, session_key, auth_token(191))
);