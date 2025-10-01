package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("memory before start test")

	// Запуск stress для нагрузки на память
	stress := exec.Command("stress", "-q", "--vm", "1")
	stress.Start()

	// Информация о процессах
	exec.Command("top", "-b", "-n", "1").Run()

	fmt.Println("memory after start test")

	// Мониторинг памяти
	exec.Command("vmstat", "2", "10").Run()

	// Сохранение ошибок ядра
	exec.Command("sh", "-c", "dmesg -eT --level=err > dmesg_errors").Run()

	// Последняя ошибка
	cmd := exec.Command("tail", "-1", "dmesg_errors")
	output, _ := cmd.Output()
	fmt.Printf("Last error: %s", output)

	// Остановка stress
	stress.Process.Kill()
}
