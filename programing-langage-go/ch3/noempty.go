package main

import "fmt"

func main() {
	tests := []string{"Math", "Science", "", "Physics"}
	tests = noempty(tests)
	fmt.Printf("%s", tests)

	//foo1 := make(map[string] int)

	foo2 := map[string]int{
		"foo":  1,
		"piyo": 2,
	}

	fmt.Printf("%s", foo2)

}

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
