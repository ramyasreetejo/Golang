package main

import (
	"fmt"
)

var j int = 1

func main() {
	fmt.Println("Permutations of a string!!")
	str := "LIFE"
	permute(str, 0)
}

func permute(str string, ind int) {
	if ind == len(str) {
		fmt.Println(j, ": ", str)
		j++
		return
	}
	for i := ind; i < len(str); i++ {
		r := []rune(str)
		r[ind], r[i] = r[i], r[ind]
		permute(string(r), ind+1)
		r[ind], r[i] = r[i], r[ind]
	}
}
