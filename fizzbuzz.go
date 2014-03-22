package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: fizzbuzz <positive int>\n")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v is not an int\n", os.Args[1])
		os.Exit(1)
	}
	if n <= 0 {
		fmt.Fprintf(os.Stderr, "Please input a positive int, not %d\n", n)
		os.Exit(1)
	}
	for i := 1; i <= n; i++ {
		var result bytes.Buffer
		result.WriteString(fmt.Sprintf("%d ", i))
		if i%3 == 0 {
			result.WriteString("fizz")
		}
		if i%5 == 0 {
			result.WriteString("buzz")
		}
		fmt.Println(result.String())
	}
}
