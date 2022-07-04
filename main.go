package main

import (
	"fmt"
	"strconv"
)

func remove(intSlice []int) []int {
	di := make(map[int]bool)
	var list []int
	for _, v := range intSlice {
		if _, val := di[v]; !val {
			di[v] = true
			list = append(list, v)
		}
	}
	return list
}
func palindrome(ff int) bool {
	var aa string
	aa = strconv.Itoa(ff)
	gg := ""
	for i := len(aa) - 1; 0 <= i; i-- {
		gg = gg + string(aa[i])
	}
	fmt.Println(gg)
	if gg == aa {
		return true
	}
	return false
}

func rpete(ff []int) []int {
	asd := []int{}
	for i := 0; i < len(ff); i++ {
		for j := 0; j < len(ff); j++ {
			if ff[i] == ff[j] && i != j {
				asd = append(asd, i)
				break
			}
		}
	}
	return asd
}
func main() {
	num := []int{1, 2, 3, 3, 9, 3, 4, 4, 5, 5, 9}
	fmt.Println(rpete(num))
}
