// iota_test
package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a, b, c)
}
