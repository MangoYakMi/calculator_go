package main

import (
	"fmt"
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	return num1 / num2
}

func calculator() {
	var num1, num2 float64
	var operator string

	fmt.Println("Добро пожаловать в калькулятор!")
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("Некорректная операция")
		return
	}

	fmt.Println("Результат:", result)
}

func main() {
	calculator()
}
