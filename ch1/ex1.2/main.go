// Echo4 prints command-line argument index and value pairs one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
