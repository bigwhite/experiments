-- name: CreateDepartment :execresult
INSERT INTO department (
  name
) VALUES (
  ?
);

-- name: GetDepartments :many
SELECT id, name FROM department;

-- name: CreateInstructor :execresult
INSERT INTO instructor (
  name, dept_id
) VALUES (
  ?, ?
);

-- name: GetInstructors :many
SELECT id, name, dept_id FROM instructor;

-- name: CreateCourse :execresult
INSERT INTO course (
  title, dept_id
) VALUES (
  ?, ?
);

-- name: GetCoursesByDept :many
SELECT id, title, dept_id FROM course WHERE dept_id = ?;

-- name: CreateStudent :execresult
INSERT INTO student (
  name, dept_id
) VALUES (
  ?, ?
);

-- name: GetStudents :many
SELECT id, name, dept_id FROM student;

-- name: EnrollStudent :execresult
INSERT INTO enrollment (
  student_id, course_id, semester, year
) VALUES (
  ?, ?, ?, ?
);

-- name: GetEnrollmentByStudent :many
SELECT student_id, course_id, semester, year FROM enrollment WHERE student_id = ?;

-- name: UpdateStudentName :exec
UPDATE student SET name = ?
WHERE id = ?;

-- name: UpdateCourseTitle :exec
UPDATE course SET title = ?
WHERE id = ?;

-- name: DeleteStudent :exec
DELETE FROM student 
WHERE id = ?;

-- name: DeleteCourse :exec
DELETE FROM course 
WHERE id = ?;

-- name: DeleteEnrollmentByCourseID :exec
DELETE FROM enrollment 
WHERE course_id = ?;
