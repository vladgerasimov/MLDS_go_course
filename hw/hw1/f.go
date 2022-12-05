//Дан текст. Выведите слово, которое в этом тексте встречается чаще всего. Если таких слов несколько,
//выведите то, которое меньше в лексикографическом порядке.
//
//Формат ввода
//Вводится текст.
//
//Формат вывода
//Выведите ответ на задачу.

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]int, 10000)
	sentence, _ := os.ReadFile("input.txt")

	words := strings.Fields(strings.TrimSuffix(string(sentence), "\n"))
	keys := make([]string, 0, len(counts))

	for _, word := range words {
		_, keyExists := counts[word]
		if keyExists == false {
			counts[word] = 0
			keys = append(keys, word)
		}
		counts[word]++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if counts[keys[i]] != counts[keys[j]] {
			return counts[keys[i]] > counts[keys[j]]
		}
		return keys[i] < keys[j]
	})

	fmt.Println(keys[0])
}
