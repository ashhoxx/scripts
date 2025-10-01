package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type LogEntry struct {
	IP    string `json:"ip"`
	Login string `json:"login"`
	Date  string `json:"date"`
}

type SSHLog struct {
	Successful struct {
		Count  int        `json:"count"`
		Logins []LogEntry `json:"logins"`
	} `json:"successful"`
	Unsuccessful struct {
		Count    int        `json:"count"`
		Attempts []LogEntry `json:"attempts"`
	} `json:"unsuccessful"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <hours>")
		return
	}

	hours := os.Args[1]

	// Получаем успешные авторизации
	successfulCmd := exec.Command("sh", "-c",
		fmt.Sprintf("journalctl -u sshd -S '%s hours ago' | grep 'Accepted'", hours))
	successfulOut, _ := successfulCmd.Output()
	successfulLines := strings.Split(string(successfulOut), "\n")

	// Получаем неуспешные попытки
	unsuccessfulCmd := exec.Command("sh", "-c",
		fmt.Sprintf("journalctl -u sshd -S '%s hours ago' | grep 'Invalid user'", hours))
	unsuccessfulOut, _ := unsuccessfulCmd.Output()
	unsuccessfulLines := strings.Split(string(unsuccessfulOut), "\n")

	var log SSHLog

	// Обрабатываем успешные авторизации
	log.Successful.Count = len(successfulLines) - 1
	for _, line := range successfulLines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 11 {
			entry := LogEntry{
				IP:    fields[10],
				Login: fields[8],
				Date:  strings.Join(fields[0:3], " "),
			}
			log.Successful.Logins = append(log.Successful.Logins, entry)
		}
	}

	// Обрабатываем неуспешные попытки
	log.Unsuccessful.Count = len(unsuccessfulLines) - 1
	for _, line := range unsuccessfulLines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 10 {
			entry := LogEntry{
				IP:    fields[9],
				Login: fields[7],
				Date:  strings.Join(fields[0:3], " "),
			}
			log.Unsuccessful.Attempts = append(log.Unsuccessful.Attempts, entry)
		}
	}

	// Сохраняем в JSON
	file, _ := os.Create("log.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(log)

	fmt.Println("Log saved to log.json")
}
