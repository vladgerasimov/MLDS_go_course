//Во входной строке записана последовательность чисел через пробел. Для каждого числа выведите
//слово YES (в отдельной строке), если это число ранее встречалось в последовательности или NO, если не встречалось.
//
//Формат ввода
//Вводится список чисел. Все числа списка находятся на одной строке.
//
//Формат вывода
//Выведите ответ на задачу.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	nums := make(map[string]int, 100000)
	input, _ := os.ReadFile("input.txt")
	numbersInput := strings.Fields(strings.TrimSuffix(string(input), "\n"))

	for _, num := range numbersInput {
		_, numExists := nums[num]
		if numExists {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
			nums[num] = 1
		}
	}
}
