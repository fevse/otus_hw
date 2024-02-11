package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		var evs EnvValue

		i, err := file.Info()
		if err != nil {
			return nil, err
		}
		if i.Size() == 0 {
			evs.NeedRemove = true
			evs.Value = ""
			env[file.Name()] = evs
			continue
		}
		pathFile := fmt.Sprintf("%s/%s", dir, file.Name())
		f, err := os.Open(pathFile)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		fl, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		ev := strings.TrimRight(string(bytes.ReplaceAll(fl, []byte{0}, []byte{10})), " ")
		evs.Value = ev
		evs.NeedRemove = true
		env[file.Name()] = evs
	}
	return env, nil
}
