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

INSERT INTO `alumnos` (`name`, `subject`, `grade`)
VALUES (`Alumno 1`, `Programacion`, `90`);

INSERT INTO `alumnos` (`name`, `subject`, `grade`)
VALUES (`Alumno 2`, `Metodos Matematicos I`, `95`);

INSERT INTO `alumnos` (`name`, `subject`, `grade`)
VALUES (`Alumno 3`, `Sistemas Distribuidos`, `98`);

/*INSERT INTO `users` (`name`) VALUES ('Isaias'), ('Aldo Jesé'), ('Itzel'), ('Victor Ulises');*/