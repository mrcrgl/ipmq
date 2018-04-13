package cli

import (
	"log"

	"github.com/mrcrgl/ipmq/src/server"
)

func Run(log *log.Logger, options Options) error {

	handler := server.New(log, "tcp", ":51892")

	if err := handler.Listen(); err != nil {
		log.Fatal(err)
	}

	// this should not be reached
	return nil
}
