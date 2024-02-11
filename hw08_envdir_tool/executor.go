package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := cmd[0]
	args := cmd[1:]
	c := exec.Command(command, args...)
	c.Env = os.Environ()
	for k, v := range env {
		str := fmt.Sprintf("%s=%s", k, v.Value)
		c.Env = append(c.Env, str)
	}
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
	return c.ProcessState.ExitCode()
}
