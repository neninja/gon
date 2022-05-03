package main

import (
	"bufio"
	"flag"
	"log"
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Fatalf("Erro: %s", err)
		}

		_, _ = buf.ReadBytes('\n')
	}
}
