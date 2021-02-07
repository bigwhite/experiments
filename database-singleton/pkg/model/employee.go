package model

type Employee struct {
	Name   string `gorm:"column:name"`
	Age    int    `gorm:"column:age"`
	Gender string `gorm:"column:gender"`
	Birth  string `gorm:"column:birthday"`
	Email  string `gorm:"column:email"`
}

func (r Employee) TableName() string {
	return "employee"
}
