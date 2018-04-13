package main

import (
	"os"

	"fmt"

	"log"

	"github.com/mrcrgl/ipmq/src/commands/ipmqd/cli"
)

func main() {
	options := cli.DefaultOptions()
	fs := cli.FlagSet(options)

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err.Error())
	}

	if err := options.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "invalid option: %s\n", err.Error())
		fs.Usage()
		os.Exit(1)
	}

	panic(cli.Run(log.New(os.Stdout, cli.Name, log.LstdFlags), *options))
}
