package prog

import (
	"bufio"
	"fmt"
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
func ReadArgs() (*ipc.Proc, rnet.Port, *crypto.Symmetric, error) {
	return readArgs(os.Args)
}

func readArgs(args []string) (proc *ipc.Proc, pool rnet.Port, key *crypto.Symmetric, err error) {
	if len(args) < 4 {
		err = ErrBadArgs
		return
	}
	log.Info(log.Lbl("ReadArgs:"), args[1], args[2], args[3])
	ipcPort, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	key, err = crypto.SymmetricFromString(args[3])
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

var stdinReader *bufio.Reader

// ReadStdin and split it up into a slice of strings, splitting at either white
// space or quotes
func ReadStdin(prompt string) []string {
	if stdinReader == nil {
		stdinReader = bufio.NewReader(os.Stdin)
	}
	fmt.Print(prompt)
	text, err := stdinReader.ReadString('\n')
	if log.Error(err) {
		return nil
	}
	return Split(text)
}

// Split a string by white space or by text surrounded by quotes. Any character
// preceded by a backspace will be rendered directly.
func Split(str string) []string {
	var strs []string
	var s []rune
	rs := []rune(str)
	ln := len(rs) - 1
	d1, d2, q1, q2 := ' ', '\t', '\'', '"'
	for i := 0; i <= ln; i++ {
		flush := i == ln
		switch r := rs[i]; r {
		case q1, q2:
			d1, d2, q1, q2 = r, 0, 0, 0
			flush = true
		case d1, d2:
			d1, d2, q1, q2 = ' ', '\t', '\'', '"'
			flush = true
		case '\\':
			if i < ln {
				s = append(s, rs[i+1])
			}
			i++
		default:
			s = append(s, r)
		}
		if flush && len(s) > 0 {
			strs = append(strs, string(s))
			s = s[0:0]
		}
	}
	return strs
}
