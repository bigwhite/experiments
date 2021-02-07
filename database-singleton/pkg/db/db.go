package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/bigwhite/testdboper/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var once sync.Once

type database struct {
	instance    *gorm.DB
	maxIdle     int
	maxOpen     int
	maxLifetime time.Duration
}

type Option func(db *database)

var db *database

func WithMaxIdle(maxIdle int) Option {
	return func(d *database) {
		d.maxIdle = maxIdle
	}
}
func WithMaxOpen(maxOpen int) Option {
	return func(d *database) {
		d.maxOpen = maxOpen
	}
}

func DB(opts ...Option) *gorm.DB {
	once.Do(func() {
		db = new(database)
		for _, f := range opts {
			f(db)
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.Config.Database.User,
			config.Config.Database.Password,
			config.Config.Database.IP,
			config.Config.Database.Port,
			config.Config.Database.DB)
		var err error
		db.instance, err = gorm.Open("mysql", dsn) // database: *gorm.DB
		if err != nil {
			panic(err)
		}

		sqlDB := db.instance.DB()
		if err != nil {
			panic(err)
		}

		if db.maxIdle != 0 {
			sqlDB.SetMaxIdleConns(db.maxIdle)
		}

		if db.maxLifetime != 0 {
			sqlDB.SetConnMaxLifetime(db.maxLifetime)
		}

		if db.maxOpen != 0 {
			sqlDB.SetMaxOpenConns(db.maxOpen)
		}

	})
	return db.instance
}
