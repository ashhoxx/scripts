package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <cpu_count>")
		return
	}

	// Начальная нагрузка
	cmd := exec.Command("uptime")
	output, _ := cmd.Output()
	fmt.Printf("load average before start: %s", output)

	// Запуск нагрузки
	stress := exec.Command("stress", "--cpu", os.Args[1])
	stress.Start()

	// Информация о процессах
	exec.Command("top", "-b", "-n", "1").Run()

	// Динамическая нагрузка
	exec.Command("vmstat", "2", "5").Run()

	// Конечная нагрузка
	cmd = exec.Command("uptime")
	output, _ = cmd.Output()
	fmt.Printf("load average after start: %s", output)

	// Остановка stress
	stress.Process.Kill()
}
