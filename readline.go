package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func readline(msg string) (line string) {
	fmt.Print(msg)
	if scanner.Scan() {
		line = scanner.Text()
	}
	return line
}
