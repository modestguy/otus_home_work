package main

import (
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

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	_, err := setEnvironments(env)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	execCommand := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	out := io.MultiWriter(os.Stdout)
	execCommand.Stderr, execCommand.Stdout = out, out

	procIn, err := execCommand.StdinPipe()
	if nil != err {
		fmt.Println(err.Error())
		return -1
	}

	go func() {
		_, err := io.Copy(io.MultiWriter(procIn), os.Stdin)
		if err != nil {
			return
		}
		err = procIn.Close()
		if err != nil {
			return
		}
	}()

	if err := execCommand.Start(); nil != err {
		fmt.Println(err.Error())
		return execCommand.ProcessState.ExitCode()
	}

	err = execCommand.Wait()
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return 0
}
