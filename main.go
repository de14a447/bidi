package main

import "fmt"

func main() {
	str, mask := "Hello, World!‮10x‭", 0

	bits := 0
	for _, char := range str {
		for char > 0 {
			bits += int(char) & mask
			char = char >> 1
		}
	}
	fmt.Println("Total bits set:", bits)
}
