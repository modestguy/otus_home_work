package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func setEnvironments(env Environment) (bool, error) {
	for k, v := range env {
		err := os.Unsetenv(k)
		if err != nil {
			return false, err
		}
		if !v.NeedRemove {
			err := os.Setenv(k, v.Value)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

var setEnvironmentsFunc = setEnvironments

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	_, err := setEnvironmentsFunc(env)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	execCommand := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	execCommand.Env = os.Environ()
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	execCommand.Stdout = mw
	execCommand.Stderr = mw

	if err := execCommand.Run(); err != nil {
		fmt.Println(err.Error())
		return -1
	}

	print(stdBuffer.String())
	return 0
}
