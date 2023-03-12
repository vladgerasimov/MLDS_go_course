package main

import (
	"bufio"
	"fmt"
	"github.com/mailru/easyjson"
	"hw3_bench/user_struct"
	"io"
	"os"
	"regexp"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("@")

	seenBrowsers := make(map[string]bool)
	uniqueBrowsers := 0
	scanner := bufio.NewScanner(file)
	fmt.Fprintln(out, "found users:")
	i := 0
	for scanner.Scan() {
		user := &user_struct.User{}
		err := easyjson.Unmarshal([]byte(scanner.Text()), user)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers
		if browsers == nil {
			i++
			continue
		}

		for _, browser := range browsers {
			if ok := strings.Contains(browser, "Android"); ok == true {
				isAndroid = true
				if _, seenBefore := seenBrowsers[browser]; !seenBefore {
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			} else if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				if _, seenBefore := seenBrowsers[browser]; !seenBefore {
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			}
		}
		if !(isAndroid && isMSIE) {
			i++
			continue
		}
		email := r.ReplaceAllString(user.Email, " [at] ")
		fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", i, user.Name, email))
		i++
	}
	//foundUsers = append(foundUsers, "")
	//fmt.Fprintln(out, strings.Join(foundUsers, "\n"))
	fmt.Fprintln(out, "\nTotal unique browsers", uniqueBrowsers)
}
