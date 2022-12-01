package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("test exec command", func(t *testing.T) {
		cmd := []string{"ls", "-lh"}
		m := make(Environment)
		code := RunCmd(cmd, m)
		require.Equal(t, 0, code)
	})

	t.Run("test executing cat command", func(t *testing.T) {
		cmd := []string{"cat"}
		m := make(Environment)
		code := RunCmd(cmd, m)
		require.Equal(t, 0, code)
	})
}
