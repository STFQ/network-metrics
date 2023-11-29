package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func executeScript() string {
	// 假设您的脚本保存在同级目录的script.sh文件中
	// 使用exec.Command执行脚本
	cmd := exec.Command("/bin/bash", "script.sh")
	stdout, err := cmd.Output()

	if err != nil {
		log.Printf("Error running script: %s", err)
		return "# ERROR: Script execution failed.\n"
	}

	return string(stdout)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && strings.Trim(r.URL.Path, "/") == "metrics" {
		// 设置HTTP头部，指定返回的内容类型为普通文本
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")

		// 执行脚本并获取输出结果
		metrics := executeScript()

		// 将结果写入HTTP响应体
		fmt.Fprint(w, metrics)
	} else {
		// 如果请求的方法不是GET或路径不是/metrics，返回404
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)

	// 监听8000端口
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
