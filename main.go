/*
Copyright © 2023 Alan Hanafy
*/
package main

import (
	"go.uber.org/zap"

	"github.com/ahanafy/promote-cli/cmd"
)

func main() {
	logger, _ := zap.NewDevelopment()

	defer func() {
		_ = logger.Sync()
	}()

	cmd.Execute()
}
