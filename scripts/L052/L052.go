package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Проверяем количество аргументов командной строки
	if len(os.Args) < 3 {
		fmt.Println("Use: --login=username [--no-password|--remove]")
		return
	}

	// Извлекаем логин пользователя из второго аргумента
	login := os.Args[2]

	// Создаем файл конфигурации sudo для пользователя в /etc/sudoers.d/
	exec.Command("touch", "/etc/sudoers.d/"+login).Run()

	// Обрабатываем команду в зависимости от первого аргумента
	switch os.Args[1] {
	case "--login":
		// Добавляем пользователя в группу sudo
		fmt.Println("add privelegies")
		exec.Command("usermod", "-a", "-G", "sudo", login).Run()

	case "--no-password":
		// Настраиваем sudo без пароля для пользователя
		fmt.Println("disable request password")
		// Добавляем правило NOPASSWD в файл конфигурации sudo
		cmd := fmt.Sprintf("echo '%s ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers.d/%s", login, login)
		exec.Command("sh", "-c", cmd).Run()
		// Показываем добавленное правило для подтверждения
		exec.Command("tail", "-n1", "/etc/sudoers.d/"+login).Run()

	case "--remove":
		// Удаляем привилегии sudo у пользователя
		fmt.Println("remove privilegies")
		// Удаляем пользователя из группы sudo
		exec.Command("deluser", login, "sudo").Run()
		// Удаляем последнее правило из файла конфигурации sudo
		exec.Command("sed", "-i", "$d", "/etc/sudoers.d/"+login).Run()
	}
}
