package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple test open file", func(t *testing.T) {
		ev, err := openFile("testdata/env/BAR")
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, "bar", ev)
	})
	t.Run("simple test read dir", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		if err != nil {
			log.Fatal(err)
		}
		require.Equal(t, 5, len(env))
	})
}
