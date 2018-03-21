/*
  UTF8 strings in Go

  In Go, all strings are UTF8 encoded.
  Each character in these strings is of type "rune"

  "Rune" is an alias for int32
  This is because Unicode chars can be 1, 2, or 4 bytes long
*/

package main

import (
	"fmt"
)

func main() {
	str := "百度一下，你就知道"
	str2 := "Hello world"

	fmt.Println(str)
	fmt.Println("String length", len([]rune(str)))
	fmt.Println("Byte length", len(str))

	fmt.Println(str2)
	fmt.Println("String length", len([]rune(str2)))
	fmt.Println("Byte length", len(str2))
}
