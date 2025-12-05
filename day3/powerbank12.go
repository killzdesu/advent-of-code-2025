package main

import (
	"fmt"
	// "strings"
	// "strconv"
	"bufio"
	"os"
)

func main() {
	answer := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		Str := []int{}
		for _, x := range line {
			Str = append(Str, int(x-'0'))
		}
		// fmt.Printf("%v", Str)
		var Max [12]int
		for i:=0;i<12;i++ {
			Max[i] = Str[len(Str)-12+i]
		}
		for i:=len(Str)-1-12;i>=0;i-- {
			pick := Str[i]
			for k:=0; k<12; k++ {
				if pick >= Max[k] {
					Max[k], pick = pick, Max[k]
				} else {
					break
				}
			}
		}
		sum := int64(0)
		for k:=0; k<12; k++ {
			sum *= 10
			sum += int64(Max[k])
		}
		answer += sum
	}
	fmt.Println(answer)
}
