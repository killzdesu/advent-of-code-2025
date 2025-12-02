package main

import (
	// "bufio"
	"fmt"
	"strconv"
)

func main() {
	var command string
	count := 0
	dial := 50

	for true {
		fmt.Scanln(&command) 
		if command == "END" {
			break
		}
		val, _ := strconv.Atoi(command[1:]) 
		if command[0] == 'L' {
			if val >= 100 {
				count += val/100
				val = val % 100
			}
			if dial >= val {
				dial -= val
			} else {
				if dial != 0 {
					count += 1
				}
				dial -= val - 100
			}
			val = 0
			if dial == 0 {
				count += 1
			}
		} else {
			dial += val
			for dial > 99 {
				dial -= 100
				count += 1
			}
		}
	}
	fmt.Printf("%d\n", count)
}
