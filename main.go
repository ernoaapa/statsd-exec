package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/cactus/go-statsd-client/statsd"
	"github.com/kelseyhightower/envconfig"
)

// StatsdConfig for statsd client
type StatsdConfig struct {
	Host   string `default:"localhost"`
	Port   int    `default:"8125"`
	Prefix string `required:"true"`
}

func main() {
	config := getConfig()
	client := getClient(config)
	defer client.Close()

	client.Inc("executed", 1, 1.0)
	start := time.Now()
	exitCode, err := run(os.Args[1:]...)
	duration := time.Now().Sub(start)

	if err != nil {
		log.Panicf("Error when resolving exit code: %v", err)
	}

	reportStats(client, exitCode, duration)

	os.Exit(exitCode)
}

func getConfig() *StatsdConfig {
	config := &StatsdConfig{}
	envconfig.MustProcess("statsd", config)
	return config
}

func getClient(config *StatsdConfig) statsd.Statter {
	client, err := statsd.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), config.Prefix)
	if err != nil {
		log.Fatal(err)
	}
	return client
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

func reportStats(client statsd.Statter, exitCode int, duration time.Duration) {
	client.TimingDuration("duration", duration, 1.0)
	if exitCode != 0 {
		client.Inc("success", 0, 1.0)
		client.Inc("failed", 1, 1.0)
	} else {
		client.Inc("success", 1, 1.0)
		client.Inc("failed", 0, 1.0)
	}
}
