package main

import (
	"github.com/julien040/anyquery/rpc"
)

func main() {
	plugin := rpc.NewPlugin(peopleCreator, passwordCreator, internetCreator)
	plugin.Serve()
}
