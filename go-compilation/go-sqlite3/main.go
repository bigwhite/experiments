package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 打开数据库（如果不存在，则创建）
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	sqlStmt := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// 插入数据
	_, err = db.Exec(`INSERT INTO user (name) VALUES (?)`, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query(`SELECT id, name FROM user;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", id, name)
	}

	// 检查查询中的错误
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
