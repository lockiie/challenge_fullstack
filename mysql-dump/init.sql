CREATE DATABASE  IF NOT EXISTS challenge_fullstack;
USE challenge_fullstack;

DROP TABLE IF EXISTS levels;

CREATE TABLE levels (
  lvl_id int NOT NULL AUTO_INCREMENT,
  lvl_name varchar(30) NOT NULL,
  PRIMARY KEY (lvl_id)
);



DROP TABLE IF EXISTS developers;

CREATE TABLE developers (
  dev_id int NOT NULL AUTO_INCREMENT,
  dev_name varchar(50) NOT NULL,
  dev_gender char(1) NOT NULL,
  dev_birth date NOT NULL,
  dev_hobby varchar(200) DEFAULT NULL,
  lvl_id int DEFAULT NULL,
  PRIMARY KEY (dev_id),
  KEY dev_lvl_fk (lvl_id),
  CONSTRAINT dev_lvl_fk FOREIGN KEY (lvl_id) REFERENCES levels (lvl_id)
);


INSERT INTO levels(lvl_name)VALUES('Beginner');
INSERT INTO levels(lvl_name)VALUES('Experient');
INSERT INTO levels(lvl_name)VALUES('Boss');

INSERT INTO developers(dev_name, dev_gender, dev_birth, dev_hobby, lvl_id)
     VALUES('Martin Fowler', 'M', STR_TO_DATE('1963-12-18', '%Y-%m-%d'), 'Read books', 3);

COMMIT;