package main

import (
	"fmt"
	"strings"
	"strconv"
)

func log10(x int) int {
	cnt := 1
	for x >= 10 {
		x /= 10
		cnt += 1
	}
	return cnt
}

func mult(x int) int {
	y := 10
	for x > 1 {
		x -= 1
		y *= 10
	}
	return y+1
}

func main() {
	var input string

	fmt.Scanln(&input)
	maxv := 0

	split_id := strings.Split(input, ",")
	data := [][]int{}
	for _, el := range split_id {
		parts := strings.Split(el, "-")
		res := []int{}
		for _, p := range parts {
			i, _ := strconv.Atoi(p)
			if i > maxv {
				maxv = i
			}
			res = append(res, i)
		}
		data = append(data, res)
	}

	ans := int64(0)
	for i := 1; i< 99999; i++ {
		y := int64(mult(log10(i)))
		x := int64(i)*y
			for _, el := range data {
				if int64(el[0]) <= x && x <= int64(el[1]) {
					ans += int64(x)
				}
			}
	}
	fmt.Println(ans)
}
