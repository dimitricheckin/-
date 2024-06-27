package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    // Чтение ввода пользователя из консоли
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Input:\n")
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Output:\nОшибка чтения ввода")
        return
    }

    // Удаление лишних пробелов вокруг введенной строки
    input = strings.TrimSpace(input)

    // Разделение строки на операнды и оператор
    operands := strings.Split(input, " ")
    if len(operands) != 3 {
        if len(operands) == 1 {
            fmt.Println("Output:\nВыдача паники, так как строка не является математической операцией.")
        } else {
            fmt.Println("Output:\nВыдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
        }
        return
    }

    var a, b int
    var operator string

    // Определение системы счисления (арабская или римская)
    isRoman := isRomanNumeral(operands[0]) && isRomanNumeral(operands[2])
    if isRomanNumeral(operands[0]) != isRomanNumeral(operands[2]) {
        fmt.Println("Output:\nВыдача паники, так как используются одновременно разные системы счисления.")
        return
    }

    // Парсинг и валидация первого операнда
    a, err = parseOperand(operands[0], isRoman)
    if err != nil {
        fmt.Println("Output:\n", err.Error())
        return
    }

    // Парсинг и валидация второго операнда
    b, err = parseOperand(operands[2], isRoman)
    if err != nil {
        fmt.Println("Output:\n", err.Error())
        return
    }

    // Валидация оператора
    operator = operands[1]
    if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
        fmt.Println("Output:\nОшибка: недопустимый оператор")
        return
    }

    // Выполнение арифметической операции
    result := calculate(a, b, operator)

    // Конвертация результата в римские цифры, если исходные числа были римскими
    if isRoman {
        if result <= 0 {
            fmt.Println("Output:\nВыдача паники, так как в римской системе нет отрицательных чисел.")
            return
        }
        fmt.Printf("Output:\n%s\n", intToRoman(result))
    } else {
        fmt.Printf("Output:\n%d\n", result)
    }
}

// Функция для парсинга и валидации операнда
func parseOperand(s string, isRoman bool) (int, error) {
    if isRoman {
        return romanToInt(s)
    } else {
        num, err := strconv.Atoi(s)
        if err != nil {
            return 0, fmt.Errorf("Ошибка: недопустимый операнд: %s", s)
        }
        if num < 1 || num > 10 {
            return 0, fmt.Errorf("Ошибка: операнд должен быть числом от 1 до 10 включительно")
        }
        return num, nil
    }
}

// Функция для выполнения арифметической операции
func calculate(a, b int, operator string) int {
    switch operator {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b == 0 {
            fmt.Println("Output:\nОшибка: деление на ноль")
            os.Exit(1)
        }
        return a / b
    default:
        fmt.Println("Output:\nОшибка: недопустимый оператор")
        os.Exit(1)
        return 0
    }
}

// Функция для проверки, является ли строка римским числом
func isRomanNumeral(s string) bool {
    switch s {
    case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
        return true
    default:
        return false
    }
}

// Функция для конвертации римского числа в арабское
func romanToInt(s string) (int, error) {
    roman := map[string]int{
        "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
        "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
    }
    if val, exists := roman[s]; exists {
        return val, nil
    }
    return 0, fmt.Errorf("Ошибка: недопустимый римский операнд: %s", s)
}

// Функция для конвертации арабского числа в римское
func intToRoman(num int) string {
    val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    var result strings.Builder
    for i := 0; i < len(val); i++ {
        for num >= val[i] {
            num -= val[i]
            result.WriteString(syb[i])
        }
    }
    return result.String()
}
