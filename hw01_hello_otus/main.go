package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func reverse(s string) string {
	s = stringutil.Reverse(s)
	return s
}

func main() {
	fmt.Println(reverse("Hello, OTUS!"))
}
