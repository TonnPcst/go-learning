package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := "ສະບາຍດີ"
	fmt.Println("A:", len(s), utf8.RuneCountInString(s))

	fmt.Println("B const:", 0.1+0.2)
	a, b := 0.1, 0.2
	fmt.Println("B vars: ", a+b)

	n := 65
	fmt.Println("C:", string(rune(n)), strconv.Itoa(n))
}
