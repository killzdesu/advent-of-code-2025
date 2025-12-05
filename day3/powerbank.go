package main

import (
	"fmt"
	// "strings"
	// "strconv"
	"bufio"
	"os"
)

func main() {
	answer := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		Str := []int{}
		for _, x := range line {
			Str = append(Str, int(x-'0'))
		}
		// fmt.Printf("%v", Str)
		Max := 0
		MaxLoc := -1
		SecMax := 0
		for in, el := range Str {
			if Max < el && in != len(Str)-1 {
				Max, MaxLoc = el, in
			}
		}
		for i := MaxLoc+1; i < len(Str); i++ {
			SecMax = max(SecMax, Str[i])
		}
		answer += 10*Max + SecMax
	}
	fmt.Println(answer)
}
