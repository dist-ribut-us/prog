package prog

import (
	"github.com/dist-ribut-us/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadArgs(t *testing.T) {
	log.To(nil)
	_, _, _, err := readArgs([]string{"notEnoughArgs"})
	assert.Equal(t, ErrBadArgs, err)

	args := []string{"command", "24369", "3000", "y58Bei1OdOjpk2DJxiQ1RdUtIwMcUbhQo9_FMum7eig="}
	proc, pool, key, err := readArgs(args)

	assert.Equal(t, 24369, int(proc.Port()))
	assert.Equal(t, 3000, int(pool))
	assert.Equal(t, args[3], key.String())
}
