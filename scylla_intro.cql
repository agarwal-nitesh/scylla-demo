CREATE KEYSPACE dev_game WITH replication = {'class':'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE users(
  id VARCHAR,
  name VARCHAR,
  email VARCHAR,
  about VARCHAR,
  PRIMARY KEY (id)
  );
Describe keyspaces;

USE dev_test;
DESCRIBE TABLE users;



INSERT INTO users(id, name, email, about) VALUES ('user1', 'foo', 'foo@example.com', 'no');
SELECT * FROM users;
ALTER TABLE users ADD type varchar ;



UPDATE users SET type= 'admin' WHERE id= 'user1';
--  INSERT does the same thing

SELECT * FROM users WHERE type ='admin';
-- Doesn't work
-- Remember Cluster ring

CREATE INDEX ON users(type);

SELECT * FROM users where id='user1' AND type='admin';
-- Doesn't work
-- Filter
