package main

import "fmt"
import "os"
import "os/exec"

const (
	path = `C:\Users\yasu\Dropbox\note\`
	cmd  = `gvim`
)

var (
	args = []string{"."}
)

func main() {
	if err := os.Chdir(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command(cmd, args...)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
