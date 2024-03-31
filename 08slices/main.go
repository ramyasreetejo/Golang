package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exploring Slices!!")

	var veggies = []string{"potato", "ladysfinger"}
	veggies = append(veggies, "onion", "brinjal")
	fmt.Println(veggies)

	veggies = append(veggies[1:]) // [:3], [2:4] etc... 2 included, until less than 4th index, 4th index excluded
	fmt.Println(veggies)

	//veggies = append(veggies, veggies[:2]) ..this is wrong syntax
	veggies = veggies[:2]
	fmt.Println(veggies)

	scores := make([]int, 3)
	scores[0] = 90
	scores[1] = 20
	scores[2] = 87
	//scores[3] = 35

	scores = append(scores, 79, 99, 100)
	fmt.Println(scores)

	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	//how to remove element from a slice w given index
	var courses = []string{"cpp", "golang", "java", "python"}
	fmt.Println(courses)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter index to be removed:")
	input, _ := reader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	// ... syntax is used to convey append function that plis accept more arguments than just what is meant for else, it will give error.
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	//2-d slices in golang

	var rows int
	var cols int

	rows = 7
	cols = 9
	var twodslices = make([][]int, rows)

	var i int

	for i = range twodslices {

		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices)
}
