package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lenth := utf8.RuneCountInString(text)

	asd := []rune(text)
	if unicode.IsUpper(asd[0]) && asd[lenth-1] == 10 {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
	fmt.Println(asd[lenth-1])
}
