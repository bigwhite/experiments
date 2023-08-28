package bar

import (
	"fmt"

	"go.uber.org/zap"
)

func Add(a, b int) int {
	fmt.Println("invoke foo.Add")
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infof("invoke bar.Add\n")
	return a + b
}
