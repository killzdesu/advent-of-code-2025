package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Pair struct {
	fi, se int64
}

func main() {
	ans := int64(0)
	scanner := bufio.NewScanner(os.Stdin)
	p := []Pair{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		a, _ := strconv.ParseInt(parts[0], 10, 64)
		b, _ := strconv.ParseInt(parts[1], 10, 64)
		p = append(p, Pair{a, b})
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].fi == p[j].fi {
			return p[i].se < p[j].se
		}
		return p[i].fi < p[j].fi
	})

	// for _, x := range p {
	// 	fmt.Printf("%v\n", x)
	// }

	for i, x := range p {
		ans += max(0, x.se-x.fi+1)
		for j:=i+1; j<len(p); j++ {
			p[j].fi = max(p[j].fi, x.se+1)
		}
	}


	for scanner.Scan() {
		// line := scanner.Text()
	}
	fmt.Println(ans)
}
