// jupyter is a web-server for Jupyter kernels
package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var (
	app *commander.Command
)

func init() {
	app = &commander.Command{
		UsageLine:   "jupyter",
		Subcommands: nil,
		Flag:        *flag.NewFlagSet("jupyter", flag.ExitOnError),
	}
}

func main() {

}
