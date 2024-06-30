package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:4407)/example_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database successfully!")

	insertData(db)
	queryData(db)
	updateData(db)
	queryData(db) // 查看更新后的数据
	deleteData(db)
	queryData(db) // 查看删除后的数据
}

func insertData(db *sql.DB) {
	// 插入department数据
	_, err := db.Exec("INSERT INTO department (name) VALUES ('Computer Science'), ('Mathematics')")
	if err != nil {
		log.Fatal(err)
	}

	// 插入instructor数据
	_, err = db.Exec("INSERT INTO instructor (name, dept_id) VALUES ('John Doe', 1), ('Jane Smith', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入course数据
	_, err = db.Exec("INSERT INTO course (title, dept_id) VALUES ('Database Systems', 1), ('Calculus', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入student数据
	_, err = db.Exec("INSERT INTO student (name, dept_id) VALUES ('Alice', 1), ('Bob', 2)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入enrollment数据
	_, err = db.Exec("INSERT INTO enrollment (student_id, course_id, semester, year) VALUES (1, 1, 'Fall', 2024), (2, 2, 'Fall', 2024)")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}

func queryData(db *sql.DB) {
	// 查询所有学生的信息
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var studentID int
		var name string
		var deptID int
		err := rows.Scan(&studentID, &name, &deptID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Student ID: %d, Name: %s, Department ID: %d\n", studentID, name, deptID)
	}

	// 查询某个院系的课程信息
	rows, err = db.Query("SELECT * FROM course WHERE dept_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var courseID int
		var title string
		var deptID int
		err := rows.Scan(&courseID, &title, &deptID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Course ID: %d, Title: %s, Department ID: %d\n", courseID, title, deptID)
	}

	// 查询某个学生的选课信息
	rows,

		err = db.Query("SELECT * FROM enrollment WHERE student_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var studentID int
		var courseID int
		var semester string
		var year int
		err := rows.Scan(&studentID, &courseID, &semester, &year)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Student ID: %d, Course ID: %d, Semester: %s, Year: %d\n", studentID, courseID, semester, year)
	}
}

func updateData(db *sql.DB) {
	// 更新某个学生的姓名
	_, err := db.Exec("UPDATE student SET name = 'Alice Johnson' WHERE student_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	// 更新某个课程的标题
	_, err = db.Exec("UPDATE course SET title = 'Advanced Database Systems' WHERE course_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data updated successfully!")
}

func deleteData(db *sql.DB) {
	// 删除某个学生的选课记录
	_, err := db.Exec("DELETE FROM enrollment WHERE student_id = ? AND course_id = ? AND semester = ? AND year = ?", 1, 1, "Fall", 2024)
	if err != nil {
		log.Fatal(err)
	}

	// 删除某个课程
	_, err = db.Exec("DELETE FROM course WHERE course_id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data deleted successfully!")
}
