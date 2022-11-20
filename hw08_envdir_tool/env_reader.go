package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	newLine      = 10
	equalsSymbol = "="
)

var ErrContainsEquals = errors.New("filename contains equal symbol")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func checkFileNameIsValid(fileName string) (bool, error) {
	res := strings.Contains(fileName, equalsSymbol)
	if res {
		return false, ErrContainsEquals
	}
	return true, nil
}

func processLine(line string) string {
	line = strings.TrimRight(line, " \t\n")
	return strings.ReplaceAll(line, "\x00", "\n")
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	resultMap := make(Environment)
	for _, file := range files {
		fileName := file.Name()
		isValid, _ := checkFileNameIsValid(fileName)
		if !isValid {
			continue
		}

		filePath := dir + "/" + fileName
		f, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		fi, err := f.Stat()
		if err != nil {
			return nil, err
		}
		if fi.Size() == 0 {
			resultMap[fileName] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
			continue
		}

		reader := bufio.NewReader(f)
		line, err := reader.ReadString(newLine)
		if err != nil && !errors.Is(io.EOF, err) {
			return nil, err
		}

		emptyLine := false
		if len(line) == 0 {
			emptyLine = true
		}

		resultMap[fileName] = EnvValue{
			Value:      processLine(line),
			NeedRemove: emptyLine,
		}

		err = f.Close()
		if err != nil {
			return nil, err
		}
	}
	return resultMap, nil
}
