package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	testEnv := Environment{
		"ONE":        {Value: "one", NeedRemove: false},
		"TO_DELETE":  {Value: "", NeedRemove: true},
		"WITH_SPACE": {Value: "some value", NeedRemove: false},
	}

	t.Run("correct reading", func(t *testing.T) {
		cmd := []string{"/bin/bash", "testdata/echo.sh", "arg1=1", "arg2=2"}
		exitCode := RunCmd(cmd, testEnv)
		require.Equal(t, 0, exitCode)
	})

	t.Run("correct reading", func(t *testing.T) {
		cmd := []string{"/bin/bash", "invalid_command_execute"}
		exitCode := RunCmd(cmd, testEnv)
		require.Equal(t, 127, exitCode)
	})
}
