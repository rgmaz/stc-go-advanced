package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positives", a: 2, b: 3, want: 5},
		{name: "with zero", a: 0, b: 7, want: 7},
		{name: "with negative", a: -2, b: -3, want: -5},
	}
	for _, c := range cases {
		got := add(c.a, c.b)
		if got == c.want {
			fmt.Printf("%s: PASS\n", c.name)
		} else {
			fmt.Printf("%s: FAIL got=%v want=%v\n", c.name, got, c.want)
		}
	}
}
