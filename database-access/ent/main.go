package main

import (
	"context"
	"log"

	"demo/ent"
	"demo/ent/course"
	"demo/ent/department"
	"demo/ent/enrollment"
	"demo/ent/student"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:4407)/example_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// 执行CRUD操作
	createData(ctx, client)
	queryData(ctx, client)
	updateData(ctx, client)
	deleteData(ctx, client)
}

func createData(ctx context.Context, client *ent.Client) {
	// 创建部门
	cs, err := client.Department.Create().SetName("Computer Science").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	math, err := client.Department.Create().SetName("Mathematics").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 创建教师
	_, err = client.Instructor.Create().SetName("John Doe").SetDepartment(cs).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Instructor.Create().SetName("Jane Smith").SetDepartment(math).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 创建课程
	dbCourse, err := client.Course.Create().SetTitle("Database Systems").SetDepartment(cs).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	calcCourse, err := client.Course.Create().SetTitle("Calculus").SetDepartment(math).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 创建学生
	alice, err := client.Student.Create().SetName("Alice").SetDepartment(cs).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	bob, err := client.Student.Create().SetName("Bob").SetDepartment(math).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 学生选课
	_, err = client.Enrollment.Create().SetStudent(alice).SetCourse(dbCourse).SetSemester("Fall").SetYear(2024).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Enrollment.Create().SetStudent(bob).SetCourse(calcCourse).SetSemester("Fall").SetYear(2024).Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func queryData(ctx context.Context, client *ent.Client) {
	// 查询所有学生
	//students, err := client.Student.Query().All(ctx)
	students, err := client.Student.Query().WithDepartment().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, stu := range students {
		log.Printf("Student ID: %d, Name: %s, Department ID: %d\n", stu.ID, stu.Name, stu.Edges.Department.ID)
	}

	// 查询某个部门的课程
	courses, err := client.Course.Query().WithDepartment().Where(course.HasDepartmentWith(department.ID(1))).All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		log.Printf("Course ID: %d, Title: %s, Department ID: %d\n", course.ID, course.Title, course.Edges.Department.ID)
	}

	// 查询某个学生的选课信息
	enrollments, err := client.Enrollment.Query().WithStudent().WithCourse().Where(enrollment.HasStudentWith(student.ID(1))).All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, enrollment := range enrollments {
		log.Printf("Student ID: %d, Course ID: %d, Semester: %s, Year: %d\n", enrollment.Edges.Student.ID,
			enrollment.Edges.Course.ID, enrollment.Semester, enrollment.Year)
	}
}

func updateData(ctx context.Context, client *ent.Client) {
	// 更新学生姓名
	_, err := client.Student.UpdateOneID(1).SetName("Alice Johnson").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 更新课程标题
	_, err = client.Course.UpdateOneID(1).SetTitle("Advanced Database Systems").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteData(ctx context.Context, client *ent.Client) {
	// 删除选课记录
	_, err := client.Enrollment.Delete().Where(enrollment.HasCourseWith(course.ID(1))).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 删除课程
	err = client.Course.DeleteOneID(1).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 删除学生
	err = client.Student.DeleteOneID(1).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
