package main

import "fmt"

func main() {

	// In ancient times one byte was enough to encode all the latin characters as well as numbers and signs.
	// A byte in Go is an unsigned 8-bit integer.
	// type byte = uint8 [-127, 127]

	ascii := 'A'
	fmt.Printf("%[1]c %[1]d %[1]b \n\n", ascii) // %c char, %d base 10, %b base 2

	var ch byte = 65
	fmt.Printf("%08b %v %c \n", ch, ch, ch)

	fmt.Println("----------------------")

	// But today computers need to handle a wide range chars and signs
	// A rune in Go is a signed 32-bit integer (4 bytes)
	// type rune = int32 [0, 42949672]

	var unicode rune = '型'                        // 💗 or 🙈
	fmt.Printf("%[1]c %[1]d %[1]b \n\n", unicode) // %c char, %d base 10, %b base 2

	var s string = "型"
	s_rune := []rune(s) // convert string to rune
	s_byte := []byte(s) // convert string to byte

	// The special Unicode character 型 rune value is 22411 but it’s taking three bytes for encoding.
	fmt.Println("型 as rune:", s_rune, " as byte:", s_byte)
	fmt.Println()

	// What does [229 158 139] mean?
	fmt.Printf("%b %b %b %s \n", 229, 158, 139, string([]byte{229, 158, 139})) // convert byte to string
	fmt.Println("----------------------")
}
