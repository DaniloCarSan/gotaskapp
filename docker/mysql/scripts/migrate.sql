CREATE DATABASE IF NOT EXISTS gotaskapp;

USE gotaskapp;

DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS status;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    firstname VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,
    email VARCHAR(60) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    verified ENUM('Y','N') DEFAULT 'N',
    createAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
)ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS  status(
    id INT PRIMARY KEY AUTO_INCREMENT, 
    name VARCHAR(30) NOT NULL,
    user_id INT NOT NULL
)ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS  tasks(
    id INT PRIMARY KEY AUTO_INCREMENT, 
    description TEXT NOT NULL,
    user_id INT NOT NULL,
    status_id INT NOT NULL
)ENGINE=INNODB;

ALTER TABLE status ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE tasks 
ADD FOREIGN KEY (user_id) REFERENCES users (id),
ADD FOREIGN KEY (status_id) REFERENCES status (id);


INSERT INTO users (firstname, lastname, email, password,verified)
-- password: 123456
VALUES ("Danilo", "Santos", "danilocarsan@gmail.com", "$2a$10$oj78CETnWUZGWiC0Wy1.cuZSVmaX.UrW.gRhhHiotaoqgBPjfr1HK",'Y');

INSERT INTO status (name, user_id) 
VALUES ("To view", 1),("To Do", 1), ("Doing", 1), ("Blocked", 1), ("Done", 1), ("Canceled", 1);

INSERT INTO tasks (description, user_id, status_id)
VALUES ("Create a project", 1, 1), ("Create a database", 1, 1), ("Create a table", 1, 1), ("Create a user", 1, 1), ("Create a status", 1, 1), ("Create a task", 1, 1);