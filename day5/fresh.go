package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	ans := 0
	scanner := bufio.NewScanner(os.Stdin)
	p := [][]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		b, _ := strconv.ParseInt(parts[1], 10, 64)
		p = append(p, []int64{a, b})
	}
	// for _, x := range p {
	// 	fmt.Printf("%v\n", x)
	// }
	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.ParseInt(line, 10, 60)
		fresh := false
		for _, x := range p {
			if x[0] <= val && val <= x[1] {
				fresh = true
			}
		}
		if fresh {
			ans++
		}
	}
	fmt.Println(ans)
}
