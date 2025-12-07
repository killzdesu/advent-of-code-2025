package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
)

func main() {
	// ans1 is for subproblem 1
	ans1 := 0
	// ans2 is for subproblem 2
	ans2 := int64(0)

	scanner := bufio.NewScanner(os.Stdin)
	s := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, []rune(line))
	}
	
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s[0]))
	}

	for j := range s[0] {
		if s[0][j] == 'S' {
			dp[0][j] = 1
		}
	}

	for i, l := range s {
		if i == len(s)-1 {
			break
		}
		for j := range l {
			if s[i][j] == 'S' && s[i+1][j] == '^' {
				if j!=0 {
					s[i+1][j-1] = 'S'
					dp[i+1][j-1] += dp[i][j]
				}
				if j!=len(l)-1 {
					s[i+1][j+1] = 'S'
					dp[i+1][j+1] += dp[i][j]
				}
				ans1++
			} 
			if s[i][j] == 'S' && s[i+1][j] != '^' {
				s[i+1][j] = 'S'
				dp[i+1][j] += dp[i][j]
			}
		}
	}

	for _, v := range (dp[len(dp)-1]) {
		ans2 += int64(v)
	}

	fmt.Println(ans2)
}
