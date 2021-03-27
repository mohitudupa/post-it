CREATE DATABASE `postit`;

CREATE TABLE `postit`.`notes` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `body` VARCHAR(1024) NOT NULL,
  `tags` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));

INSERT INTO `postit`.`notes` (`title`, `body`, `tags`) VALUES ("Hello","How ya' doing?","hey");
INSERT INTO `postit`.`notes` (`title`, `body`, `tags`) VALUES ("Heya","I'm doing well!","hey");