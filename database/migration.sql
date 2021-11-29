CREATE TABLE `users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `alumnos`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    subject varchar(255) NOT NULL,
    grade varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `alumnos` (`name`, `subject`, `grade`) VALUES ('Omar', 'Programacion', '90');

INSERT INTO `alumnos` (`name`, `subject`, `grade`) VALUES ('Isaias', 'Metodos Matematicos I', '95');

INSERT INTO `alumnos` (`name`, `subject`, `grade`) VALUES ('Carlos', 'Sistemas Distribuidos', '98');
