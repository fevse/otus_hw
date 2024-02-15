package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough args")
	}
	dir := args[1]
	env, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var cmd []string
	cmd = append(cmd, args[2:]...)
	RunCmd(cmd, env)
}
