DROP DATABASE IF EXISTS example_db;
CREATE DATABASE example_db;

-- 删除 enrollment 表
DROP TABLE enrollment;

-- 删除 student 表
DROP TABLE student; 

-- 删除 course 表
DROP TABLE course;

-- 删除 instructor 表
DROP TABLE instructor;

-- 删除 department 表
DROP TABLE department;

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
