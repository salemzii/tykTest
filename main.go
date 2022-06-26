package main

import (
	"github.com/salemzii/tykTest/dbs"
	"github.com/salemzii/tykTest/files"
	"github.com/salemzii/tykTest/logger"
)

func main() {

	dataLs := files.Reader()

	dbs.WriteData(dataLs)
	logger.InfoLogger("Successfully completed write to databases")
}
