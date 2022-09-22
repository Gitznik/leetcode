package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	l := len(s)
	var rw string
	for i := l - 1; i >= 0; i-- {
		rw = rw + string(s[i])
	}
	return rw
}

func reverseWords(s string) string {
	fmt.Println("Working on ", s)
	words := strings.Split(s, " ")
	var sdrow []string
	for _, w := range words {
		sdrow = append(sdrow, reverseWord(w))
	}
	res := strings.Join(sdrow, " ")
	fmt.Println("Result: ", res)
	return res
}

func main() {
	tst := []string{"Let's take LeetCode contest", "God Ding"}
	for _, s := range tst {
		reverseWords(s)
	}
}
