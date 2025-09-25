package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	initString := "1234567890"

	// Получаем TEST_COUNTER из переменной окружения
	testCounterStr := os.Getenv("TEST_COUNTER")
	testCounter, _ := strconv.Atoi(testCounterStr)

	if testCounter == 0 {
		// Бесконечный цикл
		counter := 1
		for {
			fmt.Printf("%s%d\n", initString, counter)
			counter++
			time.Sleep(100 * time.Millisecond) // 0.1 секунды
		}
	} else {
		// Цикл от 1 до TEST_COUNTER
		delay := time.Duration(testCounter) * 100 * time.Millisecond // TEST_COUNTER * 0.1 секунды
		for i := 1; i <= testCounter; i++ {
			fmt.Printf("%s%d\n", initString, i)
			time.Sleep(delay)
		}
	}
}
