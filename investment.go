package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
    // 1. Tạo Context để quản lý vòng đời
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // 2. Setup task chạy ngầm (VD: Server lắng nghe IoT)
    go func() {
        fmt.Println("🚀 Smart Capsule System starting...")
        // Giả lập công việc
        select {
        case <-time.After(5 * time.Second):
            fmt.Println("✅ Task completed")
        case <-ctx.Done():
            fmt.Println("🛑 Task bị hủy giữa chừng!")
        }
    }()

    // 3. Chặn hàm main lại, đợi tín hiệu từ OS (Linux/Windows)
    quit := make(chan os.Signal, 1)
    // Lắng nghe SIGINT (Ctrl+C) hoặc SIGTERM (Docker stop)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    <-quit // Block main cho đến khi nhận tín hiệu
    fmt.Println("\n⚠️  Shutting down system...")
    
    // 4. Dọn dẹp (Close DB connection, flush log...)
    cancel() 
    time.Sleep(1 * time.Second) // Đợi clean up
    fmt.Println("👋 Bye Long Đại Ca")
}