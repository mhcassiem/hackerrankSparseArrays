package main

import (
	"github.com/mhcassiem/logrus_init/logger_init"
)

func main() {
	log := logger_init.Init_logger()
	log.Info("Hello finally")
}
