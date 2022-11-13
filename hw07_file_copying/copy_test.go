package main

import (
	"errors"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("check file not exists", func(t *testing.T) {
		err := Copy("/fdgdfgdfg/fdgdfgdfg", "out.txt", 0, 0)
		require.Truef(t, errors.Is(err, syscall.ENOENT), "actual err - %v", err)
	})

	t.Run("check cannot get file size", func(t *testing.T) {
		err := Copy("/dev/urandom", "out.txt", 0, 0)
		require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual err - %v", err)
	})

	t.Run("check ofsset greater than size", func(t *testing.T) {
		err := Copy("copy.go", "out.txt", 1000000000000, 0)
		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
	})

	t.Run("check create file error", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "/cdrom/test.txt", 0, 0)
		require.Truef(t, errors.Is(err, syscall.EACCES), "actual err - %v", err)
	})

	t.Run("check seek error", func(t *testing.T) {
		err := Copy("/dev", "out.txt", 10, 0)
		require.Truef(t, errors.Is(err, syscall.EISDIR), "actual err - %v", err)
	})
}
