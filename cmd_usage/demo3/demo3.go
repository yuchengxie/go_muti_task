package main

import (
	"golang.org/x/net/context"
	"os/exec"
	"time"
)

func main() {
	var(
		ctx context.Context
		cancelFunc context.CancelFunc
	)
	ctx,cancelFunc=context.WithCancel(context.TODO())
	go func() {
		exec.CommandContext(ctx,"/bin/bash","-c","sleep 2;echo hello;")
	}()
	time.Sleep(1*time.Second)
	cancelFunc()
}
