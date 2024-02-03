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

	if limit == 0 {
		limit = s.Size() - offset
	} else if limit > s.Size()-offset {
		limit = s.Size() - offset
	}

	b := make([]byte, 1)

	rfile.Seek(offset, io.SeekStart)
	var n int64

	bar := pb.New(int(limit))
	bar.Start()

	reader := bar.NewProxyReader(rfile)

	for n < limit {
		read, err := reader.Read(b)
		n += int64(read)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		wfile.Write(b)
	}
	bar.Finish()
	wfile.Close()

	return nil
}
