package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	reg := regexp.MustCompile(`\w+('\w+)?`)
	phrase := `that's's kol"a`
	for _, word := range reg.FindAllString(strings.ToLower(phrase), -1) {
		fmt.Println(word)
	}
}
