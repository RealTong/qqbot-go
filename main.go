package main

import (
	"log"
	"net/http"
	"notify-bot/src/apis"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", apis.Test)
	mux.HandleFunc("/error", apis.DroneWebHook)
	mux.HandleFunc("/drone/webhook", apis.DroneWebHook)

	s := http.Server{
		Addr:           ":7777",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("服务启动失败: %s\n", err)
		return
	} else {
		log.Println("服务启动成功 服务监听在: ", s.Addr)
	} // 启动服务器
}
