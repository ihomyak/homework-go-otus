package main

import (
	"bufio"
	"fmt"
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
		return env, err
	}

	for _, file := range files {
		filename := file.Name()
		path := filepath.Join(dir, filename)

		info, err := file.Info()
		if err != nil {
			fmt.Printf("Не удалось получить информацию о файле %s: %v\n", filename, err)
			continue
		}

		if info.Size() == 0 {
			env[filename] = EnvValue{"", true}
			continue
		}

		if file.IsDir() || strings.Contains(filename, "=") {
			continue
		}

		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", path, err)
			continue
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", path, err)
			}
		}(f)

		scanner := bufio.NewScanner(f)
		if scanner.Scan() {
			value := scanner.Text()
			isRemove := len(value) == 0
			value = strings.TrimRight(value, " \r\t")
			value = strings.ReplaceAll(value, string(byte(0)), "\n")

			env[filename] = EnvValue{value, isRemove}
		}
	}

	return env, err
}
