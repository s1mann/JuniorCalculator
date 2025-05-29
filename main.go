package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Улучшенный калькулятор на Go")
	fmt.Println("Поддерживаемые операции: +, -, *, /, ^ (степень), sqrt (квадратный корень)")
	fmt.Println("Примеры: 2 + 3, 10.5 * 4, 2 ^ 8, sqrt 16")
	fmt.Println("Для выхода введите 'exit'")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Выход из калькулятора")
			break
		}

		result, err := advancedCalculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("Результат: %v\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
	}
}

func advancedCalculate(input string) (float64, error) {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return 0, fmt.Errorf("недостаточно аргументов")
	}

	// Обработка унарных операций (например, sqrt)
	if len(parts) == 2 {
		operator := parts[0]
		num, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, fmt.Errorf("неверное число: %v", err)
		}

		switch operator {
		case "sqrt":
			if num < 0 {
				return 0, fmt.Errorf("квадратный корень из отрицательного числа")
			}
			return math.Sqrt(num), nil
		default:
			return 0, fmt.Errorf("неподдерживаемая унарная операция: %s", operator)
		}
	}

	// Бинарные операции
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат ввода. Ожидается: число оператор число")
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("первое число неверное: %v", err)
	}

	operator := parts[1]

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("второе число неверное: %v", err)
	}

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return num1 / num2, nil
	case "^":
		return math.Pow(num1, num2), nil
	default:
		return 0, fmt.Errorf("неподдерживаемая операция: %s", operator)
	}
}
