package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func fakeSetEnvironments(Environment) (bool, error) {
	return false, errors.New("test error returned")
}

func TestRunCmd(t *testing.T) {
	t.Run("check set environments return error", func(t *testing.T) {
		setEnvironmentsFunc = fakeSetEnvironments
		defer func() { setEnvironmentsFunc = setEnvironments }()
		m := make(Environment)
		var cmd []string
		code := RunCmd(cmd, m)
		require.Equal(t, -1, code)
	})

	t.Run("test exec command", func(t *testing.T) {
		cmd := []string{"ls", "-lh"}
		m := make(Environment)
		code := RunCmd(cmd, m)
		require.Equal(t, 0, code)
	})
}
