DROP DATABASE IF EXISTS video_server;

CREATE DATABASE video_server;

USE video_server;

CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY ,
    login_name VARCHAR(64) UNIQUE KEY,
    pwd TEXT
);

CREATE TABLE video_info (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    author_id INT UNSIGNED,
    name TEXT,
    display_ctime TEXT,
    create_time DATETIME
);

CREATE TABLE comments (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    video_id VARCHAR(64),
    author_id INT UNSIGNED,
    content TEXT,
    time DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
    -- session_id TINYTEXT PRIMARY KEY NOT NULL,
    session_id VARCHAR(64) PRIMARY KEY NOT NULL,
    TTL TINYTEXT,
    login_name VARCHAR(64)
);


CREATE TABLE video_del_rec (
    video_id VARCHAR(64) PRIMARY KEY NOT NULL
);