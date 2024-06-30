DROP DATABASE IF EXISTS example_db;
CREATE DATABASE example_db;

CREATE TABLE department (
    dept_id INT AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (dept_id)
);

CREATE TABLE instructor (
    instr_id INT AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    dept_id INT,
    PRIMARY KEY (instr_id),
    FOREIGN KEY (dept_id) REFERENCES department(dept_id)
);

CREATE TABLE course (
    course_id INT AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    dept_id INT,
    PRIMARY KEY (course_id),
    FOREIGN KEY (dept_id) REFERENCES department(dept_id)
);

CREATE TABLE student (
    student_id INT AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    dept_id INT,
    PRIMARY KEY (student_id),
    FOREIGN KEY (dept_id) REFERENCES department(dept_id)
);

CREATE TABLE enrollment (
    student_id INT,
    course_id INT,
    semester VARCHAR(6),
    year INT,
    PRIMARY KEY (student_id, course_id, semester, year),
    FOREIGN KEY (student_id) REFERENCES student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(course_id)
);
