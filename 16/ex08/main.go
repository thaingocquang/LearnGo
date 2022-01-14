package main

import (
	"fmt"
	"time"
)

func main() {
	var pizzaToppings []string
	var departureTimes []time.Time
	var studentGraduationYears []int
	var statesOfLights []bool

	pizzaToppings = append(pizzaToppings, "pho mat", "tuong ot", "nho kho")
	departureTimes = append(departureTimes, time.Now(), time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*48))
	studentGraduationYears = append(studentGraduationYears, 2023, 2024, 2025)
	statesOfLights = append(statesOfLights, true, false, true)

	fmt.Println(pizzaToppings)
	fmt.Println(departureTimes)
	fmt.Println(studentGraduationYears)
	fmt.Println(statesOfLights)
}
