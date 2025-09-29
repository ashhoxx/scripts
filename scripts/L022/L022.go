package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Проверяем количество аргументов
	if len(os.Args) < 4 {
		fmt.Println("Использование: program <имя> <фамилия> <пароль>")
		os.Exit(1)
	}

	// Получаем аргументы
	name := os.Args[1]
	surname := os.Args[2]
	password := os.Args[3]

	// Создаем логин: первая буква имени + фамилия в нижнем регистре
	login := strings.ToLower(string(name[0]) + surname)

	// Создаем пользователя
	useraddCmd := exec.Command("sudo", "useradd",
		"-m",
		login,
		"-U",
		"-c", fmt.Sprintf("%s %s", name, surname),
		"-s", "/bin/bash",
		"-p", password)

	if err := useraddCmd.Run(); err != nil {
		fmt.Printf("Ошибка создания пользователя: %v\n", err)
		os.Exit(1)
	}

	// Создаем директорию .ssh
	sshDir := filepath.Join("/home", login, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		fmt.Printf("Ошибка создания директории .ssh: %v\n", err)
		os.Exit(1)
	}

	// Создаем файл authorized_keys
	authKeysFile := filepath.Join(sshDir, "authorized_keys")
	file, err := os.Create(authKeysFile)
	if err != nil {
		fmt.Printf("Ошибка создания файла authorized_keys: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Записываем SSH ключ из переменной окружения
	sshKey := os.Getenv("SSHPUBKEY")
	if sshKey != "" {
		if _, err := file.WriteString(sshKey + "\n"); err != nil {
			fmt.Printf("Ошибка записи SSH ключа: %v\n", err)
			os.Exit(1)
		}
	}

	// Устанавливаем правильные права доступа
	if err := os.Chmod(authKeysFile, 0600); err != nil {
		fmt.Printf("Ошибка установки прав доступа: %v\n", err)
		os.Exit(1)
	}

	// Меняем владельца файлов
	chownCmd := exec.Command("sudo", "chown", "-R", login+":"+login, sshDir)
	if err := chownCmd.Run(); err != nil {
		fmt.Printf("Ошибка смены владельца: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Пользователь %s успешно создан!\n", login)
}
