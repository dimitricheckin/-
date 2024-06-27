package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No input provided. Please enter a valid expression.")
			continue
		}

		if input == "exit" {
			fmt.Println("Exiting calculator...")
			break
		}

		fmt.Println("Output:")
		output := processInput(input)
		fmt.Println(output)
	}

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func processInput(input string) string {
	operands := strings.Split(input, " ")

	// Check if the input contains Roman numerals
	isRoman := false
	for _, operand := range operands {
		if isRomanNumeral(operand) {
			isRoman = true
			break
		}
	}

	if isRoman {
		// Process Roman numeral operation
		if len(operands) != 3 {
			return "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
		}

		if !isRomanNumeral(operands[0]) || !isRomanNumeral(operands[2]) {
			return "Выдача паники, так как используются одновременно разные системы счисления."
		}

		a := romanToInt(operands[0])
		b := romanToInt(operands[2])
		operator := operands[1]

		if operator == "-" && a < b {
			return "Выдача паники, так как в римской системе нет отрицательных чисел."
		}

		result := calculate(a, b, operator)

		// Convert result back to Roman numeral
		return intToRoman(result)
	} else {
		// Process Arabic numeral operation
		if len(operands) != 3 {
			if len(operands) == 1 {
				return "Выдача паники, так как строка не является математической операцией."
			}
			return "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
		}

		// Attempt to parse operands as integers
		a, err := strconv.Atoi(operands[0])
		if err != nil {
			return "Выдача паники, так как строка не является математической операцией."
		}

		b, err := strconv.Atoi(operands[2])
		if err != nil {
			return "Выдача паники, так как строка не является математической операцией."
		}

		operator := operands[1]
		result := calculate(a, b, operator)
		return fmt.Sprintf("%d", result)
	}
}



func isRomanNumeral(s string) bool {
	validRomanNumerals := map[string]bool{
		"I":  true,
		"II": true,
		"III": true,
		"IV": true,
		"V":  true,
		"VI": true,
		"VII": true,
		"VIII": true,
		"IX": true,
		"X":  true,
	}

	_, ok := validRomanNumerals[s]
	return ok
}

func romanToInt(s string) int {
	romanValues := map[string]int{
		"I":   1,
		"II":  2,
		"III": 3,
		"IV":  4,
		"V":   5,
		"VI":  6,
		"VII": 7,
		"VIII": 8,
		"IX":  9,
		"X":   10,
	}

	return romanValues[s]
}

func intToRoman(num int) string {
	if num <= 0 || num > 10 {
		panic("Выдача паники, так как результат работы с римскими числами должен быть от 1 до 10 включительно")
	}

	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return romanNumerals[num-1]
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		if a < b {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль запрещено")
		}
		return a / b
	default:
		panic("Недопустимый оператор")
	}
}
