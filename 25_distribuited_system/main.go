package main

import (
	"distribuited_system/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("node type required: master or worker")
	}
	nodeType := os.Args[1]
	switch nodeType {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode().Start()
	default:
		panic("invalid node type")
	}
}
