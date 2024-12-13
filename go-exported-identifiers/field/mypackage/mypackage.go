package mypackage

type myStruct struct {
	Field string // 导出的字段
}

// NewMyStruct1是一个导出的函数，返回myStruct的指针
func NewMyStruct1(value string) *myStruct {
	return &myStruct{Field: value}
}

// NewMyStruct1是一个导出的函数，返回myStruct类型变量
func NewMyStruct2(value string) myStruct {
	return myStruct{Field: value}
}
