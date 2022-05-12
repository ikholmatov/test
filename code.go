package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	timeStr, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	timeStr = strings.TrimSuffix(timeStr, "\n")
	mass := strings.Split(timeStr, string(','))
	tim, err := time.Parse("02.01.2006 15:04:05", mass[0])
	if err != nil {
		panic(err)
	}
	tim1, err := time.Parse("02.01.2006 15:04:05", mass[1])
	if err != nil {
		panic(err)
	}
	asd := fmt.Sprintln(tim.Sub(tim1))
	asd = strings.TrimSuffix(asd, "\x0a")
	dur, err := time.ParseDuration(asd)
	if err != nil {
		panic(err)
	}
	fmt.Println(dur)
}
