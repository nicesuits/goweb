DROP TABLE posts;
DROP TABLE threads;
DROP TABLE sessions;
DROP TABLE users;

CREATE DATABASE IF NOT EXISTS chitchatdb;

CREATE TABLE IF NOT EXISTS users
(
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  name STRING(255),
  email STRING(255) NOT NULL UNIQUE,
  password STRING(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions
(
  id SERIAL PRIMARY KEY,
  uuid STRING(64) NOT NULL UNIQUE,
  email STRING(255),
  user_id SERIAL NOT NULL REFERENCES users(id),
  created_at TIMESTAMP NOT NULL,
  INDEX (user_id)
);

CREATE TABLE IF NOT EXISTS threads
(
  id SERIAL PRIMARY KEY,
  uuid STRING(64) NOT NULL UNIQUE,
  topic STRING,
  user_id SERIAL NOT NULL REFERENCES users(id),
  created_at TIMESTAMP NOT NULL,
  INDEX (user_id)
);

CREATE TABLE IF NOT EXISTS posts
(
  id SERIAL PRIMARY KEY,
  uuid STRING(64) NOT NULL UNIQUE,
  body STRING,
  user_id SERIAL NOT NULL REFERENCES users(id),
  thread_id SERIAL NOT NULL REFERENCES threads(id),
  created_at TIMESTAMP NOT NULL,
  INDEX (user_id),
  INDEX (thread_id)
);

INSERT INTO chitchatdb.users
  VALUES (1,'411146e9-0365-4830-ae7a-168fef6d57ea', 'Peter','peter@test.com','password',TIMESTAMP '2016-03-26 10:10:10-05:00');
INSERT INTO chitchatdb.users
  VALUES (2,'06664431-eb7b-433a-b395-a09271d1ba13', 'Stephanie','stephanie@test.com','password',TIMESTAMP '2016-04-26 10:10:10-05:00');
INSERT INTO chitchatdb.threads
  VALUES (1,'411146e9-0365-4830-ae7a-168fef6d57ea', 'First Thread',1,TIMESTAMP '2016-03-26 10:10:10-05:00');
INSERT INTO chitchatdb.threads
  VALUES (2,'06664431-eb7b-433a-b395-a09271d1ba13', 'Second Thread',2,TIMESTAMP '2016-04-26 10:10:10-05:00');