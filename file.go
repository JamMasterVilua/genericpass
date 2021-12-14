package genericpass

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"
)

// OpenDefault opens default pgpass file, which is ~/.pgpass.
// Current homedir will be retrieved by calling user.Current
// or using $HOME on failure.
func OpenDefault(fileName string) (f *os.File, err error) {
	var homedir = os.Getenv("HOME")
	usr, err := user.Current()
	if err == nil {
		homedir = usr.HomeDir
	} else if homedir == "" {
		return
	}

	path := path.Join(homedir, fileName)

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if perm := uint32(info.Mode().Perm()); runtime.GOOS != "windows" && perm != 600 {
		fmt.Println("The permissions for .pgpass must be 600!")
		os.Exit(1)
	}

	return os.Open(path)
}
