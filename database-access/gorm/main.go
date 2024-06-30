package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Department struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

type Instructor struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Course struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Student struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	DeptID uint
	Dept   Department `gorm:"foreignKey:DeptID"`
}

type Enrollment struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	CourseID  uint
	Semester  string  `gorm:"size:50;not null"`
	Year      int     `gorm:"not null"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Course    Course  `gorm:"foreignKey:CourseID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:4407)/example_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移模式
	db.AutoMigrate(&Department{}, &Instructor{}, &Course{}, &Student{}, &Enrollment{})

	// 执行CRUD操作
	createData(db)
	queryData(db)
	updateData(db)
	deleteData(db)
}

func createData(db *gorm.DB) {
	// 创建部门
	cs := Department{Name: "Computer Science"}
	math := Department{Name: "Mathematics"}
	db.Create(&cs)
	db.Create(&math)

	// 创建教师
	db.Create(&Instructor{Name: "John Doe", DeptID: cs.ID})
	db.Create(&Instructor{Name: "Jane Smith", DeptID: math.ID})

	// 创建课程
	db.Create(&Course{Title: "Database Systems", DeptID: cs.ID})
	db.Create(&Course{Title: "Calculus", DeptID: math.ID})

	// 创建学生
	db.Create(&Student{Name: "Alice", DeptID: cs.ID})
	db.Create(&Student{Name: "Bob", DeptID: math.ID})

	// 学生选课
	db.Create(&Enrollment{StudentID: 1, CourseID: 1, Semester: "Fall", Year: 2024})
	db.Create(&Enrollment{StudentID: 2, CourseID: 2, Semester: "Fall", Year: 2024})
}

func queryData(db *gorm.DB) {
	// 查询所有学生
	var students []Student
	db.Find(&students)
	for _, student := range students {
		log.Printf("Student ID: %d, Name: %s, Department ID: %d\n", student.ID, student.Name, student.DeptID)
	}

	// 查询某个部门的课程
	var courses []Course
	db.Where("dept_id = ?", 1).Find(&courses)
	for _, course := range courses {
		log.Printf("Course ID: %d, Title: %s, Department ID: %d\n", course.ID, course.Title, course.DeptID)
	}

	// 查询某个学生的选课信息
	var enrollments []Enrollment
	db.Where("student_id = ?", 1).Find(&enrollments)
	for _, enrollment := range enrollments {
		log.Printf("Student ID: %d, Course ID: %d, Semester: %s, Year: %d\n", enrollment.StudentID, enrollment.CourseID, enrollment.Semester, enrollment.Year)
	}
}

func updateData(db *gorm.DB) {
	// 更新学生姓名
	db.Model(&Student{}).Where("id = ?", 1).Update("name", "Alice Johnson")

	// 更新课程标题
	db.Model(&Course{}).Where("id = ?", 1).Update("title", "Advanced Database Systems")
}

func deleteData(db *gorm.DB) {
	// 删除选课记录
	db.Where("course_id = ?", 1).Delete(&Enrollment{})

	// 删除课程
	db.Where("id = ?", 1).Delete(&Course{})
}
