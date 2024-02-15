package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("simple test run cmd", func(t *testing.T) {
		cmd := []string{"echo"}
		env := Environment{
			"FOO": EnvValue{"foo", true},
			"BAR": EnvValue{"bar", true},
		}
		ec := RunCmd(cmd, env)
		require.Equal(t, 0, ec)
	})
}
