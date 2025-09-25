package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: go run script.go <имя> <фамилия> <группа>")
		fmt.Println("Пример: go run script.go Ivan Petrov developers")
		return
	}

	name := os.Args[1]
	surname := os.Args[2]
	group := os.Args[3]

	// Генерация логина (первая буква имени + фамилия в нижнем регистре)
	login := strings.ToLower(string(name[0]) + surname)

	// Генерация случайного UID
	rand.Seed(time.Now().UnixNano())
	uid := 1000 + rand.Intn(9000)

	// Создание пользователя
	cmd := exec.Command("sudo", "useradd", "-m", login, "-u", fmt.Sprintf("%d", uid), "-g", group, "-c", name+" "+surname)
	cmd.Run()

	// Вывод информации
	fmt.Printf("Login: %s\n", login)
	fmt.Printf("Shell: /bin/bash\n")
	fmt.Printf("Home dir: /home/%s\n", login)
	fmt.Printf("Groups: %s %s\n", login, group)
	fmt.Printf("UID: %d\n", uid)
}
