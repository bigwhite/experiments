package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// 经理alice访问员工信息
	ok, err := e.Enforce("alice", "employee_info", "read") // role, obj, act
	if err != nil {
		panic(err)
	}
	fmt.Println("manager alice can read employee_info:", ok)

	ok, err = e.Enforce("alice", "employee_info", "write")
	if err != nil {
		panic(err)
	}
	fmt.Println("manager alice can write employee_info:", ok)

	// 员工bob访问自己信息
	ok, err = e.Enforce("bob", "employee_info", "write")
	fmt.Println("employee bob can write employee_info:", ok)

	// HR cathy 访问员工信息
	ok, err = e.Enforce("cathy", "employee_info", "write")
	fmt.Println("hr cathy can write employee_info:", ok)
	ok, err = e.Enforce("cathy", "employee_salary", "write")
	fmt.Println("hr cathy can write employee_salary:", ok)

	// 财务dan访问工资信息
	ok, err = e.Enforce("dan", "employee_salary", "read")
	fmt.Println("finance dan can read employee_salary:", ok)

	// 员工bob串改薪水信息
	ok, err = e.Enforce("bob", "employee_salary", "write")
	fmt.Println("employee bob can write employee_salary:", ok)
}
