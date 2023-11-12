package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func createTable(db *sql.DB, sqlStmt string) error {
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println("prepare statement error:", err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("exec prepared statement error:", err)
		return err
	}

	return nil
}

func createTables(db *sql.DB) error {
	// 创建Students表
	err := createTable(db, `CREATE TABLE IF NOT EXISTS Students (
    Sno INTEGER PRIMARY KEY,
    Sname TEXT,
    Gender TEXT, 
    Age INTEGER
  )`)
	if err != nil {
		fmt.Println("create table Students error:", err)
		return err
	}

	// 创建Courses表
	err = createTable(db, `CREATE TABLE IF NOT EXISTS Courses (
    Cno INTEGER PRIMARY KEY,
    Cname TEXT,
    Credit INTEGER
  )`)
	if err != nil {
		fmt.Println("create table Courses error:", err)
		return err
	}

	// 2022选课表
	err = createTable(db, `CREATE TABLE CourseSelection2022 (
  Sno INTEGER,
  Cno INTEGER,
  Score INTEGER,

  PRIMARY KEY (Sno, Cno),
  FOREIGN KEY (Sno) REFERENCES Students(Sno),
  FOREIGN KEY (Cno) REFERENCES Courses(Cno)
)`)
	if err != nil {
		fmt.Println("create table CourseSelection2022 error:", err)
		return err
	}

	// 2023选课表
	err = createTable(db, `CREATE TABLE CourseSelection2023 (
  Sno INTEGER,
  Cno INTEGER, 
  Score INTEGER,
  
  PRIMARY KEY (Sno, Cno),
  FOREIGN KEY (Sno) REFERENCES Students(Sno),
  FOREIGN KEY (Cno) REFERENCES Courses(Cno)
)`)

	if err != nil {
		fmt.Println("create table CourseSelection2023 error:", err)
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertData(db *sql.DB) {
	// 向Students表插入数据
	stmt, err := db.Prepare("INSERT INTO Students VALUES (?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(1001, "张三", "M", 20)
	checkErr(err)
	_, err = stmt.Exec(1002, "李四", "F", 18)
	checkErr(err)
	_, err = stmt.Exec(1003, "王五", "M", 19)
	checkErr(err)

	// 向Courses表插入数据
	stmt, err = db.Prepare("INSERT INTO Courses VALUES (?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(1, "数据库", 4)
	checkErr(err)
	_, err = stmt.Exec(2, "数学", 2)
	checkErr(err)
	_, err = stmt.Exec(3, "英语", 3)
	checkErr(err)

	// 插入2022选课数据
	stmt, _ = db.Prepare("INSERT INTO CourseSelection2022 VALUES (?, ?, ?)")
	_, err = stmt.Exec(1001, 1, 85)
	checkErr(err)
	_, err = stmt.Exec(1001, 2, 80)
	checkErr(err)
	_, err = stmt.Exec(1002, 1, 83)
	checkErr(err)
	_, err = stmt.Exec(1003, 1, 76)
	checkErr(err)
	// ...

	// 插入2023选课数据
	stmt, _ = db.Prepare("INSERT INTO CourseSelection2023 VALUES (?, ?, ?)")
	stmt.Exec(1001, 3, 75)
	checkErr(err)
	stmt.Exec(1002, 2, 81)
	checkErr(err)
	stmt.Exec(1003, 3, 86)
	checkErr(err)
}

func main() {
	db, err := sql.Open("sqlite", "../test.db")
	defer db.Close()
	if err != nil {
		fmt.Println("open test.db error:", err)
		return
	}

	err = createTables(db)
	if err != nil {
		fmt.Println("create table error:", err)
		return
	}

	insertData(db)
}
