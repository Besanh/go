package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reverse: %q\n", rev)
	fmt.Printf("Reverse again: %q\n", doubleRev)
}

// Lay chuoi dao nguoc
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
