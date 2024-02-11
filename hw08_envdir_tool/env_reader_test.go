package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple test read dir", func(t *testing.T) {
		env, _ := ReadDir("testdata/env")
		require.Equal(t, 5, len(env))
	})
}
