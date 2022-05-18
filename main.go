package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// THANKS: https://stackoverflow.com/a/22896706
func clear(osName string) {
	if osName == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

var Version = ""

func main() {
	useClear := flag.Bool("c", false, "clear terminal before every run")
	useVersion := flag.Bool("v", false, "version")
	flag.Parse()

	if *useVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	commands := flag.Args()
	if len(commands) == 0 {
		log.Fatal("command is required")
	}

	command := commands[0]
	args := flag.Args()[1:]

	buf := bufio.NewReader(os.Stdin)
	for {
		if *useClear {
			clear(runtime.GOOS)
		}

		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		_ = cmd.Run()

		_, _ = buf.ReadBytes('\n')
	}
}
