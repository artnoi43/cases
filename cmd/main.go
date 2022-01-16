package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	texts := os.Args[1:]
	for _, text := range texts {
		fmt.Printf("%s\n", strings.ToLower(text))
		fmt.Printf("%s\n", strings.ToUpper(text))
	}
}
