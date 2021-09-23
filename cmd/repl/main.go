package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dstdfx/scroopy/repl"
)

const scroopyASCIIName = `
  ______ ___________  ____   ____ ______ ___.__.
 /  ___// ___\_  __ \/  _ \ /  _ \\____ <   |  |
 \___ \\  \___|  | \(  <_> |  <_> )  |_> >___  |
/____  >\___  >__|   \____/ \____/|   __// ____|
     \/     \/                    |__|   \/
`

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(scroopyASCIIName)
	fmt.Printf("Hello %s! This is the Scroopy programming language!\n", currentUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
