//Дано натуральное число n>1. Выведите его наименьший делитель, отличный от 1.
//Алгоритм должен иметь сложность O(sqrt(n)).
//Указание. Если у числа n нет делителя не превосходящего sqrt(n), то число n — простое и ответом будет само число n.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(2)
	} else {
		for denumenator := 3; float64(denumenator) <= math.Sqrt(float64(n)); denumenator += 2 {
			if n%denumenator == 0 {
				fmt.Println(denumenator)
				return
			}
		}
		fmt.Println(n)

	}
}
