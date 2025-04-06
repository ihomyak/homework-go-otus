package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("correct reading", func(t *testing.T) {
		expected := Environment{
			"BAR":   {"bar", false},
			"UNSET": {"", true},
			"EMPTY": {"", false},
			"FOO":   {"   foo\nwith new line", false},
			"HELLO": {"\"hello\"", false},
		}

		actual, err := ReadDir("testdata/env")
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})

	t.Run("not existing dir", func(t *testing.T) {
		_, err := ReadDir("not_existing_dir")
		require.Error(t, err)
	})

	t.Run("empty dir", func(t *testing.T) {
		dir, err := os.MkdirTemp("", "fakedDir")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)

		fmt.Println(dir)
		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Equal(t, Environment{}, env)
	})
}
