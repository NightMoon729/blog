package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}
	db, err := initDB(cfg.Database)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()
	c := NewContainer(db)
	registerRoutes(c)
	startServer(cfg.Server.Port)
}

func startServer(port string) {
	addr := fmt.Sprintf(":%s", port)
	log.Printf("服务启动在 %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
