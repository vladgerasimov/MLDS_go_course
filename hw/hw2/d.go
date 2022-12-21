/*
В генеалогическом древе у каждого человека, кроме родоначальника, есть ровно один родитель.
Каждом элементу дерева сопоставляется целое неотрицательное число, называемое высотой.
У родоначальника высота равна 0, у любого другого элемента высота на 1 больше, чем у его родителя.
Вам дано генеалогическое древо, определите высоту всех его элементов.

Формат ввода
Программа получает на вход число элементов в генеалогическом древе N. Далее
следует N-1 строка, задающие родителя для каждого элемента древа, кроме родоначальника. Каждая строка
имеет вид имя_потомка имя_родителя.

Формат вывода
Программа должна вывести список всех элементов древа в лексикографическом порядке.
После вывода имени каждого элемента необходимо вывести его высоту.
*/

package main

import (
	"fmt"
	"sort"
)

var family = make(map[string]string)

func height(person string) int {
	prev, ok := family[person]
	if !ok {
		return 0
	}
	return 1 + height(prev)
}

func main() {
	var n int
	fmt.Scanln(&n)
	//scanner := bufio.NewScanner(os.Stdin)
	seenNames := make(map[string]bool, n)
	var keys []string
	var prevName, nextName string
	for i := 0; i < n-1; i++ {
		fmt.Scanln(&nextName, &prevName)
		family[nextName] = prevName
		if _, ok := seenNames[nextName]; !ok {
			seenNames[nextName] = true
			keys = append(keys, nextName)
		}
		if _, ok := seenNames[prevName]; !ok {
			seenNames[prevName] = true
			keys = append(keys, prevName)
		}
	}
	sort.Strings(keys)

	for _, name := range keys {
		fmt.Println(name, height(name))
	}
}
