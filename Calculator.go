package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
	_ "strings"
)

func getSplitOperator(text string) (string, []string) {
	var data []string
	op := ""

	switch {
	case strings.Contains(text, "+"):
		op = "+"
		data = strings.Split(text, "+")
		return op, data
	case strings.Contains(text, "-"):
		op = "-"
		data = strings.Split(text, "-")
		return op, data
	case strings.Contains(text, "*"):
		op = "*"
		data = strings.Split(text, "*")
		return op, data
	case strings.Contains(text, "/"):
		op = "/"
		data = strings.Split(text, "/")
		return op, data
	default:
		fmt.Printf("Ошибка - не корректный ввод математической операции")
		os.Exit(0)
	}

	return op, data
}
func ConvertToRomeNumb(n int) string {
	var result string
	if n < 1 {
		fmt.Printf("Ошибка - В римском режиме 0 и отрицательные числа отсутствуют")
		os.Exit(0)
	}
	if n == 100 {
		return "C"
	}
	for n > 0 {
		if n >= 50 {
			if n >= 90 {
				result += "XC"
				n -= 90
			} else if n >= 50 {
				result += "L"
				n -= 50
			}
		}
		if n >= 40 {
			result += "XL"
			n -= 40
		} else if n >= 10 {
			result += "X"
			n -= 10
		} else if n == 9 {
			result += "IX"
			n -= 9
		} else if n >= 5 {
			result += "V"
			n -= 5
		} else if n == 4 {
			result += "IV"
			n -= 4
		} else if n > 0 {
			result += "I"
			n -= 1
		}
	}
	return result
}
func MathMethods(a, b int, operation string) int {
	var result int
	switch {
	case operation == "+":
		result = a + b
		return result
	case operation == "-":
		result = a - b
		return result
	case operation == "*":
		result = a * b
		return result
	case operation == "/":
		result = a / b
		return result
	default:
		fmt.Printf("Ошибка - не корректный ввод математической операции")
		os.Exit(0)
	}

	return result
}

func Calculating(data []string, rome map[string]int, mathOp string) {

	left, err1 := strconv.Atoi(data[0])
	right, err2 := strconv.Atoi(data[1])
	romeLeft, err3 := rome[data[0]]
	romeRight, err4 := rome[data[1]]

	if err2 != nil && err4 == false {
		fmt.Printf("Ошибка - не корректный ввод данных")
		os.Exit(0)
	}
	if err1 != nil && err3 == false {
		fmt.Printf("Ошибка - не корректный ввод данных")
		os.Exit(0)
	}

	if err1 == nil && err2 == nil {
		if left > 10 || left < 1 || right > 10 || right < 1 {
			fmt.Printf("Ошибка - доступный диапазон ввода : числа от 1 до 10")
			os.Exit(0)
		}
	}
	if (err4 == true && err3 == false) ||
		(err3 == true && err4 == false) {
		fmt.Printf("Ошибка - Не корректный ввод данных. Арабско-римский ввод чисел не поддерживается. Доступный диапазон от 1 до 10")
		os.Exit(0)
	}
	if err3 == false && err4 == false || err1 != nil && err2 != nil {
		if left > 0 && right > 0 {
		fmt.Printf(strconv.Itoa(MathMethods(left, right, mathOp)) + "\n")
			}
	}
	if err3 == true && err4 == true {
		localResult := MathMethods(romeLeft, romeRight, mathOp)
		fmt.Printf(ConvertToRomeNumb(localResult) + "\n")
	}

}

func main() {

	rome := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Введите значение: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")
		op, data := getSplitOperator(text)
		if len(data) > 2 {
			fmt.Printf("Ошибка - доступно только два операнда и одна мат. операция")
			os.Exit(0)
		}
		Calculating(data, rome, op)
	}
}
