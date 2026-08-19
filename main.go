package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrTooSmall = errors.New("value too small")

func validate(n int) error {
	if n >= 10 {
		return nil
	}

	return fmt.Errorf("validating n=%d: %w", n, ErrTooSmall)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	err := validate(n)
	if err != nil {
		if errors.Is(err, ErrTooSmall) {
			fmt.Printf("too small: %d", n)
			return
		}
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println("ok")
}
