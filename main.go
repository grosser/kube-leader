package main

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"syscall"
)

var log logr.Logger

func main() {
	setupLogging()

	// wait here until we become the leader
	if err := Become(context.TODO(), os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("kube-leader: %v", err.Error()))
		os.Exit(1)
	}

	binary, err := exec.LookPath(os.Args[2])
	if err != nil {
		panic(err)
	}

	err = syscall.Exec(binary, os.Args[2:], os.Environ())
	if err != nil {
		panic(err)
	}
}

// configure logging which is used by leader.go
func setupLogging() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	loggerConfig.EncoderConfig.MessageKey = "message"
	loggerConfig.DisableCaller = true
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	log = zapr.NewLogger(logger)
}
