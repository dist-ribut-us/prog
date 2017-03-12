package prog

import (
	"github.com/dist-ribut-us/crypto"
	"github.com/dist-ribut-us/errors"
	"github.com/dist-ribut-us/ipc"
	"github.com/dist-ribut-us/log"
	"github.com/dist-ribut-us/rnet"
	"os"
	"runtime"
	"strconv"
)

var root string

// Root location for dist.ribut.us data
func Root() string {
	if root == "" {
		root = UserHomeDir() + ".ribut/"
	}
	return root
}

// ErrBadArgs is returned if there are not enough args
const ErrBadArgs = errors.String("Bad command line args")

// ReadArgs expects to be invoked with an ipcPort, the pool port and the key
// that should be used to access the merkle tree.
func ReadArgs() (*ipc.Proc, rnet.Port, *crypto.Shared, error) {
	return readArgs(os.Args)
}

func readArgs(args []string) (proc *ipc.Proc, pool rnet.Port, key *crypto.Shared, err error) {
	if len(args) < 4 {
		err = ErrBadArgs
		return
	}
	log.Info(log.Lbl("ReadArgs:"), args[1], args[2], args[3])
	ipcPort, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	key, err = crypto.SharedFromString(args[3])
	if err != nil {
		return
	}
	port, err := strconv.Atoi(args[2])
	if err != nil {
		return
	}
	pool = rnet.Port(port)
	proc, err = ipc.New(rnet.Port(ipcPort))
	return
}

// UserHomeDir get the home directory for most systems
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home + "\\"
	}
	return os.Getenv("HOME") + "/"
}
