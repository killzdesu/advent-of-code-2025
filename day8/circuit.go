package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Star struct {
	id int
	x, y, z int64
}

type pair struct {
	fi, se Star
	dist int64 
}

func getDist(a, b Star) int64 {
	return (a.x-b.x)*(a.x-b.x)+(a.y-b.y)*(a.y-b.y)+(a.z-b.z)*(a.z-b.z)
}

func getHead(h []int, x int) int {
	if h[x] == x {
		return x
	}
	h[x] = getHead(h, h[x])
	return h[x]
}

func main() {
	ans1 := int64(0)
	ans2 := int64(0)
	connection := 1000
	s := []Star{}
	d := []pair{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		ints := make([]int64, 0, len(items))
		for _, it := range items {
			if n, err := strconv.ParseInt(it, 10, 64); err == nil {
				ints = append(ints, n)
			}
		}
		t := Star{len(s), ints[0], ints[1], ints[2]}
		for _, ss := range s {
			d = append(d, pair{ss, t, int64(getDist(ss, t))})
		}
		s = append(s, t)
	}

	sort.Slice(d, func (i, j int) bool {
		return d[i].dist < d[j].dist
	})
	
	if len(d) < 1000 {
		connection = 10
	}

	h := make([]int, len(s))
	cnt := make([]int, len(s))
	for i := range h {
		h[i] = i
		cnt[i] = 1
	}

	for i:=0; i<connection; i++ {
		fi := d[i].fi.id
		se := d[i].se.id
		head := min(getHead(h, fi), getHead(h, se))
		if h[fi] == h[se] {
			continue
		}
		if getHead(h, fi) > getHead(h, se) {
			fi, se = se, fi
		}
		cnt[getHead(h, fi)] += cnt[getHead(h, se)]
		cnt[getHead(h, se)] = 0
		h[h[se]] = head
		if cnt[h[fi]] == len(s) {
			ans2 = d[i].fi.x*d[i].se.x
			break
		}
	}

	cnt1 := make([]int, len(cnt))
	copy(cnt1, cnt)
	
	sort.Slice(cnt1, func(i, j int) bool {
		return cnt1[i] > cnt1[j]
	})

	ans1 = int64(cnt1[0])*int64(cnt1[1]*cnt1[2])

	// Continue to problem 2
	for i := range d {
		fi := d[i].fi.id
		se := d[i].se.id
		head := min(getHead(h, fi), getHead(h, se))
		if h[fi] == h[se] {
			continue
		}
		if getHead(h, fi) > getHead(h, se) {
			fi, se = se, fi
		}
		cnt[getHead(h, fi)] += cnt[getHead(h, se)]
		cnt[getHead(h, se)] = 0
		h[h[se]] = head
		if cnt[h[fi]] == len(s) {
			ans2 = d[i].fi.x*d[i].se.x
			break
		}
	}
	
	fmt.Println(ans1, ans2)
}
