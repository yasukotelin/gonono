package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func newCmd(command string, r io.Reader, w io.Writer) *exec.Cmd {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd
}

func newOpenCmd(path string) *exec.Cmd {
	var command string
	switch runtime.GOOS {
	case "windows":
		command = fmt.Sprintf("start %s", path)
	case "darwin":
		// 本当にこれで開けるかは未テスト
		command = fmt.Sprintf("open %s", path)
	case "unix":
		// 本当にこれで開けるかは未テスト
		command = fmt.Sprintf("xdg-open %s", path)
	}
	return newCmd(command, nil, nil)
}

func newFinderCmd(paths []string, w io.Writer) *exec.Cmd {
	return newCmd("fzf", strings.NewReader(strings.Join(paths, "\n\r")), w)
}
