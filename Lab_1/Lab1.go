package main

import (
	"fmt"
)

func main() {

	var firstNum int
	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanf("%d\n", &firstNum)
	if err != nil {
		fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
		return
	}

	var secondNum int
	fmt.Println("Введите второе число: ")
	_, err = fmt.Scanf("%d\n", &secondNum)
	if err != nil {
		fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
		return
	}

	var operation string
	fmt.Println("Выберите операцию (+, -, *, /): ")
	_, err = fmt.Scanf("%s\n", &operation)
	if err != nil {
		fmt.Println("Ошибка ввода. Пожалуйста, используйте символы +, -, * или /")
		return
	}

	var result int

	if operation == "+" {
		result = firstNum + secondNum
		fmt.Println("Результат", result)
	} else if operation == "-" {
		result = firstNum - secondNum
		fmt.Println("Результат", result)
	} else if operation == "*" {
		result = firstNum * secondNum
		fmt.Println("Результат", result)
	} else if operation == "/" {
		if secondNum == 0 {
			fmt.Println("Делить на 0 нельзя!")
			return
		}
		result = firstNum / secondNum
		fmt.Println("Результат", result)
	} else {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
	}
}
