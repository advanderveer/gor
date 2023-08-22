// Package main provides the `gor` cli
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ())
}
