package main

import (
	"fmt"
)

// 定义安全等级
type SecurityLevel int

const (
	// 最低安全等级
	LevelLow SecurityLevel = iota
	// 中等安全等级
	LevelMedium
	// 最高安全等级
	LevelHigh
)

// 定义资源
type Resource struct {
	// 资源名称
	Name string
	// 安全等级
	Level SecurityLevel
}

// 定义用户
type User struct {
	// 用户名
	Name string
	// 安全等级
	Level SecurityLevel
}

// 定义访问控制策略
func CheckAccess(user User, resource Resource) bool {
	// 检查用户的安全等级是否高于或等于资源的安全等级
	return user.Level >= resource.Level
}

func main() {
	// 创建资源
	resource := Resource{
		Name:  "敏感数据",
		Level: LevelHigh,
	}

	// 创建用户
	user := User{
		Name:  "管理员",
		Level: LevelHigh,
	}

	// 检查访问权限
	if CheckAccess(user, resource) {
		fmt.Printf("用户[%s]有权访问资源\n", user.Name)
	} else {
		fmt.Printf("用户[%s]没有权限访问资源\n", user.Name)
	}

	// 创建用户
	user = User{
		Name:  "访客",
		Level: LevelLow,
	}

	// 检查访问权限
	if CheckAccess(user, resource) {
		fmt.Printf("用户[%s]有权访问资源\n", user.Name)
	} else {
		fmt.Printf("用户[%s]没有权限访问资源\n", user.Name)
	}
}
