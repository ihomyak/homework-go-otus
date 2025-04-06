package main

import (
	"fmt"
	"os"
)

func main() {
	// Place your code here.

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s dir program [args...]\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]
	program := os.Args[2:]

	// Читаем все файлы в директории
	env, err := ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", dir, err)
		os.Exit(1)
	}

	code := RunCmd(program, env)
	os.Exit(code)
}
