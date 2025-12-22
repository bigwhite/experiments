package main

import (
	"fmt"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func main() {
	// 初始化 go-feature-flag SDK
	// 这里我们配置它从本地文件读取规则
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second,
		Retriever: &fileretriever.Retriever{
			Path: "flags.yaml",
		},
	})
	if err != nil {
		panic(err)
	}
	// 确保程序退出时关闭 SDK，清理资源
	defer ffclient.Close()

	// 模拟当前请求的用户ID
	userID := "user-123"

	// 创建评估上下文 (Evaluation Context)
	// 这包含了判断 Flag 所需的用户信息
	userCtx := ffcontext.NewEvaluationContext(userID)

	// ❌ 痛点：
	// 代码与 "go-feature-flag" 强绑定。
	// ffclient.BoolVariation 是特定库的 API。
	// 如果未来要迁移到 LaunchDarkly 或自研系统，必须修改这里所有的调用代码。
	hasDiscount, _ := ffclient.BoolVariation("holiday-promo", userCtx, false)

	if hasDiscount {
		fmt.Printf("User %s gets a discount!\n", userID)
	} else {
		fmt.Printf("User %s pays full price.\n", userID)
	}
}
