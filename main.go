package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для определения, является ли строка числом
func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// Функция для определения является ли строка целым числом
func isFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// Функция для конвертации римских цифр в арабские
func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	total := 0
	prevValue := 0
	for _, r := range roman {
		value := romanNumerals[r]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

// Функция для конвертации арабских чисел в римские
func arabicToRoman(arabic int) string {
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	roman := ""
	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			roman += numeral.Symbol
			arabic -= numeral.Value
		}
	}
	return roman
}

// Функция для определения, является ли строка римским числом
func isRoman(str string) bool {
	romanNumerals := "IVXLCDM"
	for _, char := range str {
		if !strings.ContainsRune(romanNumerals, char) {
			return false
		}
	}
	return true
}

// Функция для выполнения операции сложения
func add(a, b int) int {
	return a + b
}

// Функция для выполнения операции вычитания
func subtract(a, b int) int {
	return a - b
}

// Функция для выполнения операции умножения
func multiply(a, b int) int {
	return a * b
}

// Функция для выполнения операции деления
func divide(a, b int) int {
	return a / b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение:")
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			panic("Неверный формат ввода")
		}

		aStr, op, bStr := parts[0], parts[1], parts[2]
		if !isNumber(aStr) || !isNumber(bStr) {
			if !isFloat(aStr) || !isFloat(bStr) {
				if isRoman(aStr) && isRoman(bStr) {
					a := romanToArabic(aStr)
					b := romanToArabic(bStr)
					if a > 10 || b > 10 {
						panic("Числа должны быть в диапазоне от 1 до 10")
					}
					result := 0
					switch op {
					case "+":
						result = add(a, b)
					case "-":
						result = subtract(a, b)
					case "*":
						result = multiply(a, b)
					case "/":
						result = divide(a, b)
					default:
						panic("Неверная операция")
					}
					if result < 1 {
						panic("Результат работы калькулятора с римскими числами может быть только положительным числом")
					} else {
						fmt.Println("Результат:", arabicToRoman(result))
						fmt.Println("Введите выражение:")
					}
				} else {
					panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
				}
			} else {
				panic("Калькулятор умеет работать только с целыми числами")
			}

		} else {
			a, _ := strconv.Atoi(aStr)
			b, _ := strconv.Atoi(bStr)
			if (a < 1 || a > 10) || (b < 1 || b > 10) {
				panic("Числа должны быть в диапазоне от 1 до 10")
			}
			result := 0
			switch op {
			case "+":
				result = add(a, b)
			case "-":
				result = subtract(a, b)
			case "*":
				result = multiply(a, b)
			case "/":
				result = divide(a, b)
			default:
				panic("Неверная операция")
			}
			fmt.Println("Результат:", result)
			fmt.Println("Введите выражение:")
		}
	}
}
