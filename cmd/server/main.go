package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"highqps-cache/internal/api"
	"highqps-cache/internal/db"
	"highqps-cache/internal/service"
)

func main() {

	// ⚠️ 修改成你的 MySQL 配置
	dsn := "root:password@tcp(127.0.0.1:3306)/test"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 测试连接（非常建议加）
	if err := conn.Ping(); err != nil {
		log.Fatal("mysql connect failed:", err)
	}

	fmt.Println("mysql connected")

	dbLayer := db.NewDB(conn)
	svc := service.NewService(dbLayer)

	http.HandleFunc("/item", api.GetItemHandler(svc))

	fmt.Println("server running at :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}