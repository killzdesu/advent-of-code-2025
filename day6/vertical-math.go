package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
)

func main() {
	ans := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	// p := [][]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	long := []int{}
	cnt := 0
	for i,_ := range lines[len(lines)-1] {
		cnt++
		if i==len(lines[0])-1 || lines[len(lines)-1][i+1]!=' ' {
			if i == len(lines[0])-1 {
				cnt++
			}
			long = append(long, cnt-1)
			cnt=0
		}
	}

	op := strings.Fields(lines[len(lines)-1])
	// fmt.Printf("%v\n", long)

	id:=0
	for ix, x := range long {
		cnt := int64(0)
		if op[ix] == "*" {
			cnt = int64(1)
		}
		for i:=id+x-1; i>=id; i-- {
			num := int64(0)
			for j:=0; j<len(lines)-1; j++ {
				if lines[j][i] != ' ' {
					num *= 10
					// val = 0
					val := lines[j][i]-'0'
					num += int64(val)
				}
			}
			if op[ix] == "*" {
				cnt *= num
			} else {
				cnt += num
			}
		}
		id += x+1
		ans += cnt
	}

	// for x, line := lines {
	// 	if f[0][0]!='*' && f[0][0]!='+' {
	// 		tmp := []int64{}
	// 		for _, x := range f {
	// 			v, _ := strconv.ParseInt(x, 10, 64)
	// 			tmp = append(tmp, v)
	// 		}
	// 		p = append(p, tmp)
	// 	} else {
	// 		break
	// 	}
	// }
	fmt.Println(ans)
}
