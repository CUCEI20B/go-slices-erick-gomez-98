package main

import "fmt"

func main() {
	var v int
	var s []int
	fmt.Scanln(&v)
	s = append(s, v)
	for i := 0; i < s[0]; i++ {
		fmt.Scanln(&v)
		s = append(s, v)
	}
	out := 0
	for _, va := range s[1:] {
		out += va
	}
	fmt.Println(out)
}
