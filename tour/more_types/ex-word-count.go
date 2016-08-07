package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var a []string = strings.Split(s, " ")
	
	for _, w := range a {
		_, present := m[w]
		if present {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
