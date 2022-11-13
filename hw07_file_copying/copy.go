package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fi, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	size := fi.Size()
	if size == 0 {
		return ErrUnsupportedFile
	}

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Seek(offset, 0)
	if err != nil {
		return err
	}
	writeFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func(writeFile *os.File) {
		err := writeFile.Close()
		if err != nil {
			panic(err)
		}
	}(writeFile)

	if limit == 0 {
		limit = size
	}

	_, err = io.CopyN(writeFile, file, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
