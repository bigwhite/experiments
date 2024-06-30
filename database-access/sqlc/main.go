package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"demo/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:4407)/example_db"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// 执行CRUD操作
	createData(queries)
	queryData(queries)
	updateData(queries)
	deleteData(queries)
}

func createData(queries *db.Queries) {
	ctx := context.Background()

	// 创建部门
	_, err := queries.CreateDepartment(ctx, "Computer Science")
	if err != nil {
		log.Fatal(err)
	}
	_, err = queries.CreateDepartment(ctx, "Mathematics")
	if err != nil {
		log.Fatal(err)
	}

	// 创建教师
	_, err = queries.CreateInstructor(ctx, db.CreateInstructorParams{Name: "John Doe", DeptID: sql.NullInt32{1, true}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = queries.CreateInstructor(ctx, db.CreateInstructorParams{Name: "Jane Smith", DeptID: sql.NullInt32{2, true}})
	if err != nil {
		log.Fatal(err)
	}

	// 创建课程
	_, err = queries.CreateCourse(ctx, db.CreateCourseParams{Title: "Database Systems", DeptID: sql.NullInt32{1, true}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = queries.CreateCourse(ctx, db.CreateCourseParams{Title: "Calculus", DeptID: sql.NullInt32{2, true}})
	if err != nil {
		log.Fatal(err)
	}

	// 创建学生
	_, err = queries.CreateStudent(ctx, db.CreateStudentParams{Name: "Alice", DeptID: sql.NullInt32{1, true}})
	if err != nil {
		log.Fatal(err)
	}
	_, err = queries.CreateStudent(ctx, db.CreateStudentParams{Name: "Bob", DeptID: sql.NullInt32{2, true}})
	if err != nil {
		log.Fatal(err)
	}

	// 学生选课
	_, err = queries.EnrollStudent(ctx, db.EnrollStudentParams{StudentID: sql.NullInt32{1, true}, CourseID: sql.NullInt32{1, true}, Semester: "Fall", Year: 2024})
	if err != nil {
		log.Fatal(err)
	}
	_, err = queries.EnrollStudent(ctx, db.EnrollStudentParams{StudentID: sql.NullInt32{2, true}, CourseID: sql.NullInt32{2, true}, Semester: "Fall", Year: 2024})
	if err != nil {
		log.Fatal(err)
	}
}

func queryData(queries *db.Queries) {
	ctx := context.Background()

	// 查询所有学生
	students, err := queries.GetStudents(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, student := range students {
		fmt.Printf("Student ID: %d, Name: %s, Department ID: %d\n", student.ID, student.Name, student.DeptID.Int32)
	}

	// 查询某个部门的课程
	courses, err := queries.GetCoursesByDept(ctx, sql.NullInt32{1, true})
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		fmt.Printf("Course ID: %d, Title: %s, Department ID: %d\n", course.ID, course.Title, course.DeptID.Int32)
	}

	// 查询某个学生的选课信息
	enrollments, err := queries.GetEnrollmentByStudent(ctx, sql.NullInt32{1, true})
	if err != nil {
		log.Fatal(err)
	}
	for _, enrollment := range enrollments {
		fmt.Printf("Student ID: %d, Course ID: %d, Semester: %s, Year: %d\n", enrollment.StudentID.Int32, enrollment.CourseID.Int32, enrollment.Semester, enrollment.Year)
	}
}

func updateData(queries *db.Queries) {
	ctx := context.Background()

	// 更新学生姓名
	err := queries.UpdateStudentName(ctx, db.UpdateStudentNameParams{ID: 1, Name: "Alice Johnson"})
	if err != nil {
		log.Fatal(err)
	}

	// 更新课程标题
	err = queries.UpdateCourseTitle(ctx, db.UpdateCourseTitleParams{ID: 1, Title: "Advanced Database Systems"})
	if err != nil {
		log.Fatal(err)
	}
}

func deleteData(queries *db.Queries) {
	ctx := context.Background()

	// 删除选课记录
	err := queries.DeleteEnrollmentByCourseID(ctx, sql.NullInt32{1, true})
	if err != nil {
		log.Fatal(err)
	}

	// 删除课程
	err = queries.DeleteCourse(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	// 删除学生
	err = queries.DeleteStudent(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
}
