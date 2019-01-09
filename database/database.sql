CREATE DATABASE dbmp;

USE dbmp;

CREATE TABLE students
(
    id int NOT NULL AUTO_INCREMENT,
    nome varchar(50) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE employees
(
    id int NOT NULL AUTO_INCREMENT,
    nome varchar(50) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO students (nome) VALUES
("Carolina"),
("Fulano"),
("Sicrano");

INSERT INTO employees (nome) VALUES
("Zé da Couve"),
("Maria Silva"),
("João Pereira");