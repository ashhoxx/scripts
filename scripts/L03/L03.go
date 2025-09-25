package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Инстуркция по использованию скрипта
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run . <PID|имя> <kill|forcekill>")
		fmt.Println("Примеры:")
		fmt.Println("  go run . 1234 kill")
		fmt.Println("  go run . myapp forcekill")
		os.Exit(1)
	}

	target := os.Args[1] //PID или имя процесса
	mode := os.Args[2]   //Определяет тип команды (kill of forcekill)

	// Определяем, является ли target числом (PID) или строкой (имя)
	isPID, _ := regexp.MatchString(`^\d+$`, target)

	//Вызов соответсвующей функции
	if isPID {
		pid, _ := strconv.Atoi(target)
		handlePID(pid, mode)
	} else {
		handleProcessName(target, mode)
	}
}

func handlePID(pid int, mode string) {
	fmt.Printf("Обработка PID: %d\n", pid)

	// Обработка ошибки
	if !processExists(pid) {
		fmt.Printf("Процесс с PID %d не существует\n", pid)
		return
	}

	//Отправка соответсвующего сигнала
	switch mode {
	case "kill":
		sendSignal(pid, "SIGTERM")
		fmt.Printf("Отправлен SIGTERM процессу %d\n", pid)
	case "forcekill":
		sendSignal(pid, "SIGTERM")
		fmt.Printf("Отправлен SIGTERM процессу %d\n", pid)

		time.Sleep(10 * time.Second)

		//processExists - проверяет существование процессов по имени
		if processExists(pid) {
			sendSignal(pid, "SIGKILL")
			fmt.Printf("Процесс %d не завершился, отправлен SIGKILL\n", pid)
		} else {
			fmt.Printf("Процесс %d успешно завершен\n", pid)
		}
	default:
		fmt.Println("Неизвестный режим. Используйте kill или forcekill")
	}
}

func handleProcessName(name string, mode string) { // Функция выполняем те же самые действия, только по имени
	fmt.Printf("Обработка процессов с именем: %s\n", name)

	if !processNameExists(name) {
		fmt.Printf("Процессы с именем '%s' не найдены\n", name)
		return
	}

	switch mode {
	case "kill":
		exec.Command("pkill", "-15", name).Run()
		fmt.Printf("Отправлен SIGTERM всем процессам '%s'\n", name)
	case "forcekill":
		exec.Command("pkill", "-15", name).Run()
		fmt.Printf("Отправлен SIGTERM всем процессам '%s'\n", name)

		time.Sleep(10 * time.Second)

		if processNameExists(name) {
			exec.Command("pkill", "-9", name).Run()
			fmt.Printf("Процессы '%s' не завершились, отправлен SIGKILL\n", name)
		} else {
			fmt.Printf("Все процессы '%s' успешно завершены\n", name)
		}
	default:
		fmt.Println("Неизвестный режим. Используйте kill или forcekill")
	}
}

// Функции проверок
func processExists(pid int) bool {
	_, err := os.FindProcess(pid)
	return err == nil
}

func processNameExists(name string) bool {
	cmd := exec.Command("pgrep", "-f", name)
	output, _ := cmd.Output()
	return len(strings.TrimSpace(string(output))) > 0
}

// sendSignal(pid int, signal string) - отправляет сигнал процессу
func sendSignal(pid int, signal string) {
	cmd := exec.Command("kill", "-"+getSignalNumber(signal), strconv.Itoa(pid))
	cmd.Run()
}

// getSignalNumber(signal string) - преобразует имя сигнала в номер
func getSignalNumber(signal string) string {
	switch signal {
	case "SIGTERM":
		return "15"
	case "SIGKILL":
		return "9"
	default:
		return "15"
	}
}
