package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:4407)/example_db"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database successfully!")

	insertData(db)
	queryData(db)
	updateData(db)
	queryData(db) // 查看更新后的数据
	deleteData(db)
	queryData(db) // 查看删除后的数据
}

func insertData(db *sqlx.DB) {
	// 插入department数据
	_, err := db.NamedExec(`INSERT INTO department (name) VALUES (:name)`, []map[string]interface{}{
		{"name": "Computer Science"},
		{"name": "Mathematics"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 插入instructor数据
	_, err = db.NamedExec(`INSERT INTO instructor (name, dept_id) VALUES (:name, :dept_id)`, []map[string]interface{}{
		{"name": "John Doe", "dept_id": 1},
		{"name": "Jane Smith", "dept_id": 2},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 插入course数据
	_, err = db.NamedExec(`INSERT INTO course (title, dept_id) VALUES (:title, :dept_id)`, []map[string]interface{}{
		{"title": "Database Systems", "dept_id": 1},
		{"title": "Calculus", "dept_id": 2},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 插入student数据
	_, err = db.NamedExec(`INSERT INTO student (name, dept_id) VALUES (:name, :dept_id)`, []map[string]interface{}{
		{"name": "Alice", "dept_id": 1},
		{"name": "Bob", "dept_id": 2},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 插入enrollment数据
	_, err = db.NamedExec(`INSERT INTO enrollment (student_id, course_id, semester, year) VALUES (:student_id, :course_id, :semester, :year)`, []map[string]interface{}{
		{"student_id": 1, "course_id": 1, "semester": "Fall", "year": 2024},
		{"student_id": 2, "course_id": 2, "semester": "Fall", "year": 2024},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}

type Student struct {
	StudentID int    `db:"student_id"`
	Name      string `db:"name"`
	DeptID    int    `db:"dept_id"`
}

type Course struct {
	CourseID int    `db:"course_id"`
	Title    string `db:"title"`
	DeptID   int    `db:"dept_id"`
}

type Enrollment struct {
	StudentID int    `db:"student_id"`
	CourseID  int    `db:"course_id"`
	Semester  string `db:"semester"`
	Year      int    `db:"year"`
}

func queryData(db *sqlx.DB) {
	// 查询所有学生的信息
	var students []Student
	err := db.Select(&students, "SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}
	for _, student := range students {
		fmt.Printf("Student ID: %d, Name: %s, Department ID: %d\n", student.StudentID, student.Name, student.DeptID)
	}

	// 查询某个院系的课程信息
	var courses []Course
	err = db.Select(&courses, "SELECT * FROM course WHERE dept_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		fmt.Printf("Course ID: %d, Title: %s, Department ID: %d\n", course.CourseID, course.Title, course.DeptID)
	}

	// 查询某个学生的选课信息
	var enrollments []Enrollment
	err = db.Select(&enrollments, "SELECT * FROM enrollment WHERE student_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	for _, enrollment := range enrollments {
		fmt.Printf("Student ID: %d, Course ID: %d, Semester: %s, Year: %d\n", enrollment.StudentID, enrollment.CourseID, enrollment.Semester, enrollment.Year)
	}
}

func updateData(db *sqlx.DB) {
	// 更新某个学生的姓名
	_, err := db.NamedExec("UPDATE student SET name = :name WHERE student_id = :student_id", map[string]interface{}{
		"name":       "Alice Johnson",
		"student_id": 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 更新某个课程的标题
	_, err = db.NamedExec("UPDATE course SET title = :title WHERE course_id = :course_id", map[string]interface{}{
		"title":     "Advanced Database Systems",
		"course_id": 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data updated successfully!")
}

func deleteData(db *sqlx.DB) {
	// 删除某个学生的选课记录
	_, err := db.NamedExec("DELETE FROM enrollment WHERE student_id = :student_id AND course_id = :course_id AND semester = :semester AND year = :year", map[string]interface{}{
		"student_id": 1,
		"course_id":  1,
		"semester":   "Fall",
		"year":       2024,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 删除某个课程
	_, err = db.NamedExec("DELETE FROM course WHERE course_id = :course_id", map[string]interface{}{
		"course_id": 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data deleted successfully!")
}
