package main

import (
	"bufio"
	"flag"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	command := flag.Args()[0]
	args := flag.Args()[1:]

	buf := bufio.NewReader(os.Stdin)
	for {
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		_ = cmd.Run()

		_, _ = buf.ReadBytes('\n')
	}
}
