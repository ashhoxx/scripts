package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run script.go <package_list_file> <install|uninstall>")
		return
	}

	packageList := os.Args[1]
	command := os.Args[2]

	file, err := os.Open(packageList)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if command == "install" {
		for scanner.Scan() {
			pkg := scanner.Text()
			exec.Command("apt", "install", pkg, "-y").Run()
		}
	} else if command == "uninstall" {
		for scanner.Scan() {
			pkg := scanner.Text()
			exec.Command("apt", "remove", pkg, "-y").Run()
		}
	}
}
