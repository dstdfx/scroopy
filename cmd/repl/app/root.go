package app

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/dstdfx/scroopy/repl"
)

// Variables that are injected in build time.
var (
	buildGitCommit string
	buildGitTag    string
	buildDate      string
	buildCompiler  = runtime.Version()
)

// TODO: add build-in function that prints build info

const scroopyASCIIName = `
  ______ ___________  ____   ____ ______ ___.__.
 /  ___// ___\_  __ \/  _ \ /  _ \\____ <   |  |
 \___ \\  \___|  | \(  <_> |  <_> )  |_> >___  |
/____  >\___  >__|   \____/ \____/|   __// ____|
     \/     \/                    |__|   \/
`

func Run() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(scroopyASCIIName)
	fmt.Printf("Hello %s! This is the Scroopy programming language!\n", currentUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
