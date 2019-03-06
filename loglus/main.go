package main

import (
	"github.com/kiwamunet/go-samples/loglus/logger"
)

var log = logger.LogNew()

func main() {
	log.Println("main starting...")

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}
