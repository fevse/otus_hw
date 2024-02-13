package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
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

		evs.NeedRemove = checkEnv(file.Name())

		if i.Size() == 0 {
			evs.Value = ""
			env[file.Name()] = evs
			continue
		}
		pathFile := filepath.Join(dir, file.Name())
		ev, err := openFile(pathFile)
		if err != nil {
			return nil, err
		}
		evs.Value = ev
		env[file.Name()] = evs
	}
	return env, nil
}

func openFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	fl, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	ev := strings.TrimRight(string(bytes.ReplaceAll(fl, []byte{0}, []byte{10})), " ") // byte{0} = NUL, byte{10} = \n
	return ev, nil
}

func checkEnv(v string) bool {
	osEnv := os.Environ()
	for _, e := range osEnv {
		if strings.Contains(e, v+"=") {
			return true
		}
	}
	return false
}
