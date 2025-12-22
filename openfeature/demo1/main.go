package main

import (
	"fmt"
	"os"
)

func main() {
	// 模拟当前请求的用户ID
	userID := "user-123"

	// ❌ 痛点：
	// 1. 全局生效：一旦开启，所有用户都会看到。
	// 2. 修改需要重启：必须修改环境变量并重启服务才能生效。
	// 3. 逻辑僵化：无法实现“只对 user-123 开启”这样的规则。

	// 从环境变量获取开关状态
	enablePromo := os.Getenv("ENABLE_HOLIDAY_PROMO") == "true"

	if enablePromo {
		fmt.Printf("User %s gets a discount!\n", userID)
	} else {
		fmt.Printf("User %s pays full price.\n", userID)
	}
}
