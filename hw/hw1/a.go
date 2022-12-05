//Напишите функцию min4(a, b, c, d), вычисляющую минимум четырех чисел, которая не содержит инструкции if, а использует стандартную функцию min. Считайте четыре целых числа и выведите их минимум.
//
//Формат ввода
//Вводятся четыре целых числа.
//
//Формат вывода
//Выведите ответ на задачу.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	maxAB := math.Min(a, b)
	maxCD := math.Min(c, d)
	fmt.Println(math.Min(maxAB, maxCD))
}
