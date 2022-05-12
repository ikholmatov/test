package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	readed, _ := read.ReadString('\n')
	sp := strings.Split(readed, string(','))
	
}
