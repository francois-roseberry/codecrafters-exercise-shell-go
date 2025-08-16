package main

import (
	"io"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	cmd := exec.Command("./shell")
	stdout, err := cmd.StdoutPipe()
	require.NoError(t, err)
	require.NoError(t, cmd.Start())
	//require.NoError(t, cmd.Wait())

	t.Run("should print prompt", func(t *testing.T) {
		read, err := io.ReadAll(stdout)
		require.NoError(t, err)
		assert.Equal(t, "$ ", string(read))
	})
}
