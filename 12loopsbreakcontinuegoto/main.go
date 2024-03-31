package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exploring loops, break, continue and goto in golang.")

	days := []string{}
	days = append(days, "Monday", "Wednesday", "Friday")

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for ind, val := range days {
		fmt.Printf("For index %v the value is %v.", ind, val)
	}

	index := 1
	for index < 10 {
		if index == 2 {
			goto label
		}
		if index == 3 || index == 6 {
			index++
			continue
		} else if index == 8 {
			break
		}
		fmt.Println(index)
		index++
	}
label:
	fmt.Println("You jumped to this label using goto.")

}
