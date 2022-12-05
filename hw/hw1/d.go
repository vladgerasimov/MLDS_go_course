//Известно, что на доске 8×8 можно расставить 8 ферзей так, чтобы они не били друг друга. Вам дана
//расстановка 8 ферзей на доске, определите, есть ли среди них пара бьющих друг друга.
//
//Формат ввода
//Программа получает на вход восемь пар чисел, каждое число от 1 до 8 - координаты 8 ферзей.
//
//Формат вывода
//Если ферзи не бьют друг друга, выведите слово NO, иначе выведите YES.

package main

import (
	"fmt"
	"math"
)

func main() {
	figs := make([][]float64, 8)
	for i := 0; i < 8; i++ {
		figs[i] = make([]float64, 2)
	}

	for i := 0; i < 8; i++ {
		fmt.Scanln(&figs[i][0], &figs[i][1])
	}

	for idxOuter, figureOuter := range figs {
		for idxInner, figureInner := range figs {
			if idxInner != idxOuter {
				i1, i2, j1, j2 := figureOuter[0], figureInner[0], figureOuter[1], figureInner[1]
				if (i1 == i2) || (j1 == j2) || math.Abs(i1-i2) == math.Abs(j1-j2) {
					fmt.Println("YES")
					return
				}
			}
		}
	}
	fmt.Println("NO")
}
