package main

import (
	"context"
	"fmt"
	"time"

	// OpenFeature SDK
	"github.com/open-feature/go-sdk/openfeature"

	// GO Feature Flag In Process Provider
	gofeatureflaginprocess "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag-in-process/pkg"

	// GO Feature Flag 配置
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func main() {
	// ==========================================
	// A. 初始化层 (Infrastructure Layer)
	// ==========================================

	ctx := context.Background()

	// 1. 创建 GO Feature Flag In Process Provider
	options := gofeatureflaginprocess.ProviderOptions{
		GOFeatureFlagConfig: &ffclient.Config{
			PollingInterval: 3 * time.Second,
			Context:         ctx,
			Retriever: &fileretriever.Retriever{
				Path: "flags.yaml",
			},
		},
	}

	provider, err := gofeatureflaginprocess.NewProviderWithContext(ctx, options)
	if err != nil {
		panic(fmt.Errorf("failed to create provider: %v", err))
	}
	defer provider.Shutdown()

	// 2. 设置 OpenFeature Provider 并等待初始化完成
	err = openfeature.SetProviderAndWait(provider)
	if err != nil {
		panic(fmt.Errorf("failed to set provider: %v", err))
	}

	fmt.Println("✅ OpenFeature In-Process provider is ready!")

	// ==========================================
	// B. 业务逻辑层 (Business Logic Layer)
	// ==========================================

	// 1. 获取 OpenFeature 客户端
	client := openfeature.NewClient("app-backend")

	// 2. 准备评估上下文
	userID := "user-123"
	evalCtx := openfeature.NewEvaluationContext(
		userID,
		map[string]interface{}{
			"email": "test@example.com",
		},
	)

	// 3. 评估 Flag
	hasDiscount, err := client.BooleanValue(
		context.Background(),
		"holiday-promo", // Flag Key
		false,           // Default Value
		evalCtx,         // Context
	)

	if err != nil {
		fmt.Printf("Error evaluating flag: %v\n", err)
	}

	if hasDiscount {
		fmt.Printf("✅ User %s gets a discount!\n", userID)
	} else {
		fmt.Printf("❌ User %s pays full price.\n", userID)
	}

	// ==========================================
	// C. 测试其他用户
	// ==========================================

	fmt.Println("\n--- Testing another user ---")

	anotherUserCtx := openfeature.NewEvaluationContext(
		"user-456",
		map[string]interface{}{
			"email": "another@example.com",
		},
	)

	hasDiscountAnother, err := client.BooleanValue(
		context.Background(),
		"holiday-promo",
		false,
		anotherUserCtx,
	)

	if err != nil {
		fmt.Printf("Error evaluating flag: %v\n", err)
	}

	if hasDiscountAnother {
		fmt.Printf("✅ User user-456 gets a discount!\n")
	} else {
		fmt.Printf("❌ User user-456 pays full price.\n")
	}

	// ==========================================
	// D. 展示更复杂的评估上下文示例
	// ==========================================

	fmt.Println("\n--- Testing with detailed user context ---")

	detailedUserCtx := openfeature.NewEvaluationContext(
		"user-789",
		map[string]interface{}{
			"firstname": "john",
			"lastname":  "doe",
			"email":     "john.doe@example.com",
			"admin":     true,
			"anonymous": false,
		},
	)

	hasDiscountDetailed, err := client.BooleanValue(
		context.Background(),
		"holiday-promo",
		false,
		detailedUserCtx,
	)

	if err != nil {
		fmt.Printf("Error evaluating flag: %v\n", err)
	}

	if hasDiscountDetailed {
		fmt.Printf("✅ User user-789 gets a discount!\n")
	} else {
		fmt.Printf("❌ User user-789 pays full price.\n")
	}
}
