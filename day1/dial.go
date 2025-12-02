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
			dial -= val
			for dial < 0 {
				dial += 100
			}
		} else {
			dial += val
			for dial > 99 {
				dial -= 100
			}
		}
		if dial == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
