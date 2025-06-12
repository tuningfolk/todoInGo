DROP TABLE IF EXISTS posts
CREATE TABLE posts(
	id 	INT AUTO_INCREMENT NOT NULL,
	title 	VARCHAR(128) NOT NULL,
	body 	VARCHAR(255),
	user_id INT  NOT NULL,
	PRIMARY KEY(`id`)
);

INSERT INTO posts
	(title,user_id)
VALUES
	("finish homework", 0);

CREATE TABLE users(
	user_id 	INT AUTO_INCREMENT 	NOT NULL,
	username 	VARCHAR(128) 		NOT NULL,
	email 		VARCHAR(255) 		NOT NULL,
	password 	CHAR(64) 		NOT NULL,
	PRIMARY KEY(`user_id`)
);

INSERT INTO users
	(username, email, password)
VALUES
	('ali', 'ali@gmail.com', 'password');

