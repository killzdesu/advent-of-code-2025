package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	ans := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	p := [][]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Fields(line)
		if f[0][0]!='*' && f[0][0]!='+' {
			tmp := []int64{}
			for _, x := range f {
				v, _ := strconv.ParseInt(x, 10, 64)
				tmp = append(tmp, v)
			}
			p = append(p, tmp)
		} else {
			for i, x := range f {
				cnt := int64(0)
				if x[0]=='+' {
					for j:=0;j<len(p);j++ {
						cnt += p[j][i]
					}
				} else {
					cnt = 1
					for j:=0;j<len(p);j++ {
						cnt *= p[j][i]
					}
				}
				ans += cnt
			}
			break
		}
	}
	fmt.Println(ans)
}
