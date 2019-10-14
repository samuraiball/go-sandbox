package main

import "fmt"

func main() {
	s := []string{"a", "b", ""}

	fmt.Printf("%d", len(s))
	fmt.Printf("%d", len(s[1:]))
	fmt.Printf("%d", cap(s[1:]))
}
