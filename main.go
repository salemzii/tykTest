package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/salemzii/tykTest/dbs"
	"github.com/salemzii/tykTest/files"
	"github.com/salemzii/tykTest/logger"
)

func main() {
	LoadDotEnv()
	dataLs := files.Reader()

	dbs.WriteData(dataLs)
	logger.InfoLogger("Successfully completed write to databases")
}

func LoadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.ErrorLogger(err)
		os.Exit(1)
	}

}
