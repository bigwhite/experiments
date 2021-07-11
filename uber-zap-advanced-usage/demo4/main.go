package main

import (
	"github.com/bigwhite/zap-usage/pkg/log"
)

func main() {
	var tops = []log.TeeOption{
		{
			Filename: "access.log",
			Ropt: log.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl log.Level) bool {
				return lvl <= log.InfoLevel
			},
		},
		{
			Filename: "error.log",
			Ropt: log.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl log.Level) bool {
				return lvl > log.InfoLevel
			},
		},
	}

	logger := log.NewTeeWithRotate(tops)
	log.ResetDefault(logger)

	for i := 0; i < 20000; i++ {
		log.Info("demo3:", log.String("app", "start ok"),
			log.Int("major version", 3))
		log.Error("demo3:", log.String("app", "crash"),
			log.Int("reason", -1))
	}

}
