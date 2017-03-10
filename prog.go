package prog

import (
	"github.com/dist-ribut-us/crypto"
	"github.com/dist-ribut-us/errors"
	"github.com/dist-ribut-us/ipc"
	"github.com/dist-ribut-us/log"
	"github.com/dist-ribut-us/rnet"
	"os"
	"strconv"
)

const ErrBadArgs = errors.String("Bad command line args")

// ReadArgs expects to be invoked with an ipcPort, the pool port and the key
// that should be used to access the merkle tree.
func ReadArgs() (proc *ipc.Proc, pool *rnet.Addr, key *crypto.Shared, err error) {
	args := os.Args
	if len(args) != 4 {
		err = ErrBadArgs
		return
	}
	log.Info(log.Lbl("ReadArgs: "), args[1], args[2], args[3])
	ipcPort, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	key, err = crypto.SharedFromString(args[3])
	if err != nil {
		return
	}
	pool, err = rnet.ResolveAddr("127.0.0.1:" + args[2])
	if err != nil {
		return
	}
	proc, err = ipc.New(rnet.Port(ipcPort))
	return
}
