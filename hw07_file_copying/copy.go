package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	mode, err := os.Lstat(fromPath)
	if err != nil {
		return err
	}
	if !mode.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	rfile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer rfile.Close()

	s, err := rfile.Stat()
	if err != nil {
		return err
	}

	if s.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	wfile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer wfile.Close()

	if limit == 0 {
		limit = s.Size() - offset
	} else if limit > s.Size()-offset {
		limit = s.Size() - offset
	}

	rfile.Seek(offset, io.SeekStart)

	bar := pb.New(int(limit))
	bar.Start()

	reader := bar.NewProxyReader(rfile)

	_, err = io.CopyN(wfile, reader, limit)
	if err != nil {
		return err
	}
	bar.Finish()

	return nil
}
