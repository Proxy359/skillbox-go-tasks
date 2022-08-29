package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			fmt.Println("выхожу из программы")
			return
		case <-time.After(1 * time.Second):
			ans := 0
			fmt.Println("Введите число")
			fmt.Scan(&ans)
			fmt.Println("Квадрат", ans, "-", ans*ans, "\n")
		}
	}
}
