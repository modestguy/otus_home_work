package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const bufSize = 1

func moveBytes(reader io.Reader, writer io.Writer, limit int) error {
	bar := pb.StartNew(limit)
	buf := make([]byte, bufSize)
	sumRead := 0
	for sumRead < limit {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := writer.Write(buf[:n]); err != nil {
			return err
		}
		sumRead += n
		bar.Increment()
		time.Sleep(1000)
	}
	bar.Finish()
	return nil
}

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

	_, err = file.Seek(offset, io.SeekStart)
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

	portion := limit
	if offset+limit > size {
		portion = size - offset
	}
	err = moveBytes(file, writeFile, int(portion))
	if err != nil {
		return err
	}
	return nil
}
