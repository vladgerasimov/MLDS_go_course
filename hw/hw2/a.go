//Как известно, в США президент выбирается не прямым голосованием, а путем двухуровневого голосования.
//Сначала проводятся выборы в каждом штате и определяется победитель выборов в данном штате.
//Затем проводятся государственные выборы: на этих выборах каждый штат имеет определенное число голосов — число
//выборщиков от этого штата. На практике, все выборщики от штата голосуют в соответствии с результами голосования
//внутри штата, то есть на заключительной стадии выборов в голосовании участвуют штаты, имеющие различное число
//голосов. Вам известно за кого проголосовал каждый штат и сколько голосов было отдано данным штатом.
//Подведите итоги выборов: для каждого из участника голосования определите число отданных за него голосов.
//
//Формат ввода
//Каждая строка входного файла содержит фамилию кандидата, за которого отдают голоса выборщики этого штата,
//затем через пробел идет количество выборщиков, отдавших голоса за этого кандидата.
//
//Формат вывода
//Выведите фамилии всех кандидатов в лексикографическом порядке, затем, через пробел, количество
//отданных за них голосов.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	counts := make(map[string]int)
	var keys []string
	var name string
	var numVotes int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		name = text[0]
		numVotes, _ = strconv.Atoi(text[1])

		_, ok := counts[name]
		if ok == false {
			counts[name] = 0
			keys = append(keys, name)
		}
		counts[name] += numVotes
	}

	sort.Strings(keys)
	for _, name := range keys {
		fmt.Println(name, counts[name])
	}
}
