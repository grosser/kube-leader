package main

import (
	"context"
	"fmt"
	"github.com/operator-framework/operator-sdk/pkg/leader"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"os"
	"os/exec"
	"syscall"
	"go.uber.org/zap"
	"github.com/go-logr/zapr"
)

func main(){
	setupLogging()

	// wait here until we become the leader
	if err := leader.Become(context.TODO(), os.Args[1]); err != nil {
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

// configure global logging which is used by operator sdk
func setupLogging() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = ""
	loggerConfig.EncoderConfig.MessageKey = "message"
	loggerConfig.DisableCaller = true
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	logf.SetLogger(zapr.NewLogger(logger))
}
