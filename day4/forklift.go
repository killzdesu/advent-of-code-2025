package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	s := [][]rune{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, []rune(line))
	}
	ans := 0
	round := 1
	for round != 0 {
		round = 0
		for i,si := range s {
			for j,_ := range si {
				if s[i][j] == '@' {
					cnt := 0
					for A:=-1;A<2;A++ {
						for B:=-1;B<2;B++ {
							if i+A<0||j+B<0||i+A>=len(s)||j+B>=len(si) {
								continue
							}
							if s[i+A][j+B] == '@'{
								cnt++
							}
						}
					}
					if cnt <=4 {
						// round ++
						ans ++
						// s[i][j] = '.'
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
