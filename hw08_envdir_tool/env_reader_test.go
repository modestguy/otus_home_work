package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const emptyDirectory = "./empty"

func TestReadDir(t *testing.T) {
	t.Run("test readdir error", func(t *testing.T) {
		_, err := ReadDir("./executor.go")
		require.Error(t, err)
	})

	t.Run("test empty directory", func(t *testing.T) {
		err := os.Mkdir(emptyDirectory, 0o755)
		if err != nil {
			return
		}
		actualEnv, err := ReadDir(emptyDirectory)
		require.NoError(t, err)
		expectedEnv := make(Environment)
		require.Equal(t, expectedEnv, actualEnv)
		err = os.Remove(emptyDirectory)
		if err != nil {
			return
		}
	})

	t.Run("check incorrect filename", func(t *testing.T) {
		err := os.Mkdir(emptyDirectory, 0o755)
		if err != nil {
			return
		}
		_, err = os.Create(emptyDirectory + "/empty=test.txt")
		if err != nil {
			return
		}
		actualEnv, err := ReadDir(emptyDirectory)
		require.NoError(t, err)
		expectedEnv := make(Environment)
		require.Equal(t, expectedEnv, actualEnv)
		err = os.RemoveAll(emptyDirectory)
		if err != nil {
			return
		}
	})

	t.Run("check empty correct file", func(t *testing.T) {
		err := os.Mkdir(emptyDirectory, 0o755)
		if err != nil {
			return
		}
		_, err = os.Create(emptyDirectory + "/empty")
		if err != nil {
			return
		}
		actualEnv, err := ReadDir(emptyDirectory)
		require.NoError(t, err)
		expectedEnv := Environment{"empty": EnvValue{
			Value:      "",
			NeedRemove: true,
		}}

		require.Equal(t, expectedEnv, actualEnv)
		err = os.RemoveAll(emptyDirectory)
		if err != nil {
			return
		}
	})

	t.Run("process file  with tabs and spaces", func(t *testing.T) {
		err := os.Mkdir(emptyDirectory, 0o755)
		if err != nil {
			return
		}
		f, err := os.Create(emptyDirectory + "/empty")
		if err != nil {
			return
		}
		_, err = f.WriteString("Hello World \t")
		if err != nil {
			fmt.Println(err)
			err := f.Close()
			if err != nil {
				return
			}
			return
		}

		actualEnv, err := ReadDir(emptyDirectory)
		require.NoError(t, err)
		expectedEnv := Environment{"empty": EnvValue{
			Value:      "Hello World",
			NeedRemove: false,
		}}

		require.Equal(t, expectedEnv, actualEnv)
		err = os.RemoveAll(emptyDirectory)
		if err != nil {
			return
		}
	})

	t.Run("test data from testdata directory", func(t *testing.T) {
		actualEnv, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		expectedEnv := Environment{
			"BAR": EnvValue{
				Value:      "bar",
				NeedRemove: false,
			},
			"EMPTY": EnvValue{
				Value:      "",
				NeedRemove: false,
			},
			"FOO": EnvValue{
				Value:      "   foo\nwith new line",
				NeedRemove: false,
			},
			"HELLO": EnvValue{
				Value:      "\"hello\"",
				NeedRemove: false,
			},
			"UNSET": EnvValue{
				Value:      "",
				NeedRemove: true,
			},
		}

		require.Equal(t, expectedEnv, actualEnv)
	})
}
