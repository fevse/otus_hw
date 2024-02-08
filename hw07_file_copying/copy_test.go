package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyErr(t *testing.T) {
	t.Run("error unsupported file", func(t *testing.T) {
		r, err := os.CreateTemp("testdata", "read")
		if err != nil {
			log.Fatal(err)
		}
		w, err := os.CreateTemp("testdata", "write")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())
		defer os.Remove(r.Name())

		Copy("testdata", w.Name(), 0, 0)
		require.Error(t, ErrUnsupportedFile)
	})
	t.Run("error offset exceeds file size", func(t *testing.T) {
		r, err := os.CreateTemp("testdata", "read")
		if err != nil {
			log.Fatal(err)
		}
		w, err := os.CreateTemp("testdata", "write")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())
		defer os.Remove(r.Name())

		Copy(r.Name(), w.Name(), 1024, 0)
		require.Error(t, ErrOffsetExceedsFileSize)
	})
}

func TestCopyData(t *testing.T) {
	t.Run("test data out offset0 limit0", func(t *testing.T) {
		w, err := os.CreateTemp("testdata", "out")
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Open("testdata/out_offset0_limit0.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())

		Copy("testdata/input.txt", w.Name(), 0, 0)
		wStat, err := w.Stat()
		if err != nil {
			log.Fatal(err)
		}
		oStat, err := o.Stat()
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, wStat.Size(), oStat.Size())
	})
	t.Run("test data out offset0 limit1000", func(t *testing.T) {
		w, err := os.CreateTemp("testdata", "out")
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Open("testdata/out_offset0_limit1000.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())

		Copy("testdata/input.txt", w.Name(), 0, 1000)
		wStat, err := w.Stat()
		if err != nil {
			log.Fatal(err)
		}
		oStat, err := o.Stat()
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, wStat.Size(), oStat.Size())
	})
	t.Run("test data out offset0 limit10000", func(t *testing.T) {
		w, err := os.CreateTemp("testdata", "out")
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Open("testdata/out_offset0_limit10000.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())

		Copy("testdata/input.txt", w.Name(), 0, 10000)
		wStat, err := w.Stat()
		if err != nil {
			log.Fatal(err)
		}
		oStat, err := o.Stat()
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, wStat.Size(), oStat.Size())
	})
}

func TestCopyDataWithOffset(t *testing.T) {
	t.Run("test data out offset100 limit1000", func(t *testing.T) {
		w, err := os.CreateTemp("testdata", "out")
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Open("testdata/out_offset100_limit1000.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())

		Copy("testdata/input.txt", w.Name(), 100, 1000)
		wStat, err := w.Stat()
		if err != nil {
			log.Fatal(err)
		}
		oStat, err := o.Stat()
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, wStat.Size(), oStat.Size())
	})
	t.Run("test data out offset6000 limit1000", func(t *testing.T) {
		w, err := os.CreateTemp("testdata", "out")
		if err != nil {
			log.Fatal(err)
		}
		o, err := os.Open("testdata/out_offset6000_limit1000.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(w.Name())

		Copy("testdata/input.txt", w.Name(), 6000, 1000)
		wStat, err := w.Stat()
		if err != nil {
			log.Fatal(err)
		}
		oStat, err := o.Stat()
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, wStat.Size(), oStat.Size())
	})
}
