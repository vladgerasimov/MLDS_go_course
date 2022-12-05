//N кеглей выставили в один ряд, занумеровав их слева направо числами от 1 до N.
//Затем по этому ряду бросили K шаров, при этом i-й шар сбил все кегли с номерами от li до ri включительно.
//Определите, какие кегли остались стоять на месте.
//
//Формат ввода
//Программа получает на вход количество кеглей N и количество бросков K.
//Далее идет K пар чисел li, ri, при этом 1≤ li≤ ri≤ N ≤ 100.
//
//Формат вывода
//Программа должна вывести последовательность из N символов, где j-й символ есть “I”, если j-я кегля
//осталась стоять, или “.”, если j-я кегля была сбита.

package main

import (
	"fmt"
)

func main() {
	var pinsNum, throws int
	fmt.Scan(&pinsNum, &throws)
	pins := make([]rune, pinsNum)
	for i := range pins {
		pins[i] = 'I'
	}

	var start, end int
	for throw := 0; throw < throws; throw++ {
		fmt.Scan(&start, &end)
		for hit := start; hit <= end; hit++ {
			pins[hit-1] = '.'
		}
	}

	fmt.Println(string(pins))
}
