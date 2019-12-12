package main

import (
	"os"

	_ "github.com/cs3org/reva/cmd/revad/runtime"
	"github.com/owncloud/ocis-hello/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
