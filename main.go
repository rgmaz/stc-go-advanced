package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FilterInts(s []int, keep func(int) bool) []int {
	var r []int
	for _, v := range s {
		if keep(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		return
	}

	var nums []int
	for _, f := range strings.Fields(sc.Text()) {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}

	evens := FilterInts(nums, func(n int) bool {
		return n%2 == 0
	})

	parts := make([]string, len(evens))
	for i, v := range evens {
		parts[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(parts, " "))
}
