package main

import (
	"fmt"
)

type Resource struct {
	Name      string
	Owner     string
	AccessMap map[string]bool
}

func (r *Resource) GrantAccess(user string) {
	r.AccessMap[user] = true
}

func (r *Resource) RevokeAccess(user string) {
	r.AccessMap[user] = false
}

func (r *Resource) CanAccess(user string) bool {
	access, exists := r.AccessMap[user]
	if !exists {
		return false
	}
	return access
}

func main() {
	// 创建一个资源
	resource := Resource{
		Name:      "example.txt",
		Owner:     "alice",
		AccessMap: make(map[string]bool),
	}

	// 授予访问权限给用户
	resource.GrantAccess("alice")
	resource.GrantAccess("bob")

	// 验证访问权限
	fmt.Println("alice can access:", resource.CanAccess("alice")) // 输出: true
	fmt.Println("bob can access:", resource.CanAccess("bob"))     // 输出: true
	fmt.Println("eve can access:", resource.CanAccess("eve"))     // 输出: false

	// 撤销访问权限
	resource.RevokeAccess("bob")

	// 验证访问权限
	fmt.Println("bob can access:", resource.CanAccess("bob")) // 输出: false
}
