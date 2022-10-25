package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strings are immutable
	s := "hello你好" // We need 5 bytes for hello and 6 bytes for 你 and 好

	// Show immutability
	//for i := 0; i < len(s); i++ {
	//	s[i] = byte(1)
	//}

	// for range loop
	for i, v := range s {
		fmt.Printf("%d: %b ", i, v)
	}
	fmt.Println()
}

func strIntConversions() {

	s := "123"
	fmt.Println(strconv.Atoi(s))
	fmt.Println(strconv.ParseInt(s, 10, 8)) // convert to base 10 and store in 8 bit (int8)

	d := 123
	fmt.Println(strconv.Itoa(d))
	fmt.Println(fmt.Sprintf("%d", d))

	strconv.ParseBool("")
}

func convertStringToBool() {

	t1, _ := strconv.ParseBool("1")
	t2, _ := strconv.ParseBool("t")
	t3, _ := strconv.ParseBool("T")
	t4, _ := strconv.ParseBool("True")
	t5, _ := strconv.ParseBool("TRUE")
	t6, _ := strconv.ParseBool("true")
	fmt.Println("true is: ", t1 && t2 && t3 && t4 && t5 && t6)

	f1, _ := strconv.ParseBool("0")
	f2, _ := strconv.ParseBool("f")
	f3, _ := strconv.ParseBool("F")
	f4, _ := strconv.ParseBool("False")
	f5, _ := strconv.ParseBool("FLASE")
	f6, _ := strconv.ParseBool("false")
	fmt.Println("false is:", f1 || f2 || f3 || f4 || f5 || f6)
}
