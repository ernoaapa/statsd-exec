package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	exitCode, err := run(os.Args[1:]...)
	if err != nil {
		log.Panic(err)
	}
	os.Exit(exitCode)
}

func run(args ...string) (int, error) {
	command := exec.Command(args[0], args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()

	return resolveExitCode(err)
}

func resolveExitCode(err error) (int, error) {
	if err != nil {
		if msg, ok := err.(*exec.ExitError); ok {
			if s, ok := msg.Sys().(syscall.WaitStatus); ok {
				return int(s.ExitStatus()), nil
			}
		}
		return 0, err
	}
	return 0, nil
}
