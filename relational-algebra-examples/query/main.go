package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	_ "modernc.org/sqlite"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Students表
type Student struct {
	Sno    int
	Sname  string
	Gender string
	Age    int
}

// Courses表
type Course struct {
	Cno    int
	Cname  string
	Credit int
}

// CourseSelection表
type CourseSelection struct {
	Sno   int
	Cno   int
	Score int
}

func doSelection(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 where score >= 80")
	var selections []CourseSelection
	for rows.Next() {
		var s CourseSelection
		rows.Scan(&s.Sno, &s.Cno, &s.Score)
		selections = append(selections, s)
	}
	fmt.Println("\nselection operation:")
	fmt.Println(selections)
}

func doProjection(db *sql.DB) {
	rows, _ := db.Query("SELECT Sno, Sname FROM Students")
	var students []Student
	for rows.Next() {
		var s Student
		rows.Scan(&s.Sno, &s.Sname)
		students = append(students, s)
	}
	fmt.Println("\nprojection operation:")
	fmt.Println(students)
}

func doCompositionOperation(db *sql.DB) {
	rows, _ := db.Query("SELECT Sno, Sname FROM Students where age >= 20")
	var students []Student
	for rows.Next() {
		var s Student
		rows.Scan(&s.Sno, &s.Sname)
		students = append(students, s)
	}
	fmt.Println("\ncomposition operation:")
	fmt.Println(students)
}

func doIntersection(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 INTERSECT SELECT * FROM CourseSelection2023")
	var selections []CourseSelection
	for rows.Next() {
		var s CourseSelection
		rows.Scan(&s.Sno, &s.Cno, &s.Score)
		selections = append(selections, s)
	}
	fmt.Println("\nintersection operation:")
	fmt.Println(selections)
}

func doUnion(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 UNION SELECT * FROM CourseSelection2023")
	var selections []CourseSelection
	for rows.Next() {
		var s CourseSelection
		rows.Scan(&s.Sno, &s.Cno, &s.Score)
		selections = append(selections, s)
	}
	fmt.Println("\nunion operation:")
	fmt.Println(selections)
}

func doDifference(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 WHERE Cno NOT IN (SELECT Cno FROM CourseSelection2023)")
	var selections []CourseSelection
	for rows.Next() {
		var s CourseSelection
		rows.Scan(&s.Sno, &s.Cno, &s.Score)
		selections = append(selections, s)
	}
	fmt.Println("\ndifference operation:")
	fmt.Println(selections)
}

// StudentCourse结果
type StudentCourse struct {
	Sno    int
	Sname  string
	Gender string
	Age    int
	Cno    int
	Cname  string
	Credit int
}

func doCartesianProduct(db *sql.DB) {
	//rows, _ := db.Query("SELECT * FROM Students CROSS JOIN Courses")
	rows, _ := db.Query("SELECT Students.*, Courses.* FROM Students, Courses")
	var selections []StudentCourse
	for rows.Next() {
		var s StudentCourse
		rows.Scan(&s.Sno, &s.Sname, &s.Gender, &s.Age, &s.Cno, &s.Cname, &s.Credit)
		selections = append(selections, s)
	}
	fmt.Println("\ncartesion-product operation:")
	fmt.Println(len(selections))
	fmt.Println(selections)
}

func doEquijoin(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 JOIN Students ON CourseSelection2022.Sno = Students.Sno")
	dumpOperationResult("Equijoin", rows)
}

func doNaturaljoin(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CourseSelection2022 NATURAL JOIN Students")
	dumpOperationResult("Naturaljoin", rows)
}

func doSemijoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM Students
WHERE EXISTS (
    SELECT *
    FROM CourseSelection2022
    WHERE Students.Sno = CourseSelection2022.Sno
)`)
	dumpOperationResult("Semijoin", rows)
}

func doThetajoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM CourseSelection2022
JOIN Students ON CourseSelection2022.Sno > Students.Sno`)
	dumpOperationResult("Thetajoin", rows)
}

func doAntijoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM Students
WHERE NOT EXISTS (
    SELECT *
    FROM CourseSelection2022
    WHERE Students.Sno = CourseSelection2022.Sno
)`)
	dumpOperationResult("Antijoin", rows)
}

func doLeftjoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM Students
LEFT JOIN CourseSelection2022 ON Students.Sno = CourseSelection2022.Sno`)
	dumpOperationResult("Leftjoin", rows)
}

func doRightjoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM Students
RIGHT JOIN CourseSelection2022 ON Students.Sno = CourseSelection2022.Sno`)
	dumpOperationResult("Rightjoin", rows)
}

func doFulljoin(db *sql.DB) {
	rows, _ := db.Query(`SELECT *
FROM Students
FULL JOIN CourseSelection2022 ON Students.Sno = CourseSelection2022.Sno`)
	dumpOperationResult("Fulljoin", rows)
}

func dumpOperationResult(operation string, rows *sql.Rows) {
	cols, _ := rows.Columns()

	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)
	defer w.Flush()
	w.Write([]byte(strings.Join(cols, "\t")))
	w.Write([]byte("\n"))

	row := make([][]byte, len(cols))
	rowPtr := make([]any, len(cols))
	for i := range row {
		rowPtr[i] = &row[i]
	}

	fmt.Printf("\n%s operation:\n", operation)
	for rows.Next() {
		rows.Scan(rowPtr...)
		w.Write(bytes.Join(row, []byte("\t")))
		w.Write([]byte("\n"))
	}
}

func main() {
	db, err := sql.Open("sqlite", "../test.db")
	defer db.Close()
	if err != nil {
		fmt.Println("open test.db error:", err)
		return
	}

	// 选择(Selection)
	doSelection(db)

	// 投影(Projection)
	doProjection(db)

	// 组合(Compostion)：选择&投影
	doCompositionOperation(db)

	// 关系交(Intersection)
	doIntersection(db)

	// 关系并(Union)
	doUnion(db)

	// 关系差(Difference)
	doDifference(db)

	// 笛卡尔积(Cartesian-Product)
	doCartesianProduct(db)

	// 连接(Join)

	// 等值连接(Equijoin)
	doEquijoin(db)

	// 自然连接(Naturaljoin)
	doNaturaljoin(db)

	// θ连接(Thetajoin)
	doThetajoin(db)

	// 半连接(Semijoin)
	doSemijoin(db)

	// 反连接(Antijoin)
	doAntijoin(db)

	// 左(外)连接(Leftjoin)
	doLeftjoin(db)

	// 右(外)连接(Rightjoin)
	doRightjoin(db)

	// 全连接(Fulljoin)
	doFulljoin(db)

}
