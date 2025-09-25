package main

import (
	"fmt"
	"time"
)

func main() {
	// Текущее время
	now := time.Now()

	// Время в разных часовых поясах
	fmt.Printf("%s: Current time in UTC time zone\n", now.UTC().Format("15:04:05"))

	ny, _ := time.LoadLocation("America/New_York")
	fmt.Printf("%s: Current time in America/New_York time zone\n", now.In(ny).Format("15:04:05"))

	tokyo, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Printf("%s: Current time in Asia/Tokyo time zone\n", now.In(tokyo).Format("15:04:05"))

	// Завтрашний день
	tomorrow := now.AddDate(0, 0, 1)
	weekday := tomorrow.Weekday().String()

	// Проверка на выходные
	if weekday == "Saturday" || weekday == "Sunday" {
		fmt.Printf("%s: Tommorow day ;-)\n", weekday)
	} else {
		fmt.Printf("%s: Tommorow day\n", weekday)
	}
}
