package main

import (
	"fmt"
	"os"

	"github.com/YakDriver/terrastory/finder"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fmt.Println("Terrastory")

	tests := finder.Find()
	if len(tests) == 0 {
		fmt.Println("No tests found")
		return 1
	}

	for _, t := range tests {
		fmt.Printf("test! %v\n", t)
	}

	return 0
}
