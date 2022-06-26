package main

import (
	"fmt"

	"github.com/salemzii/tykTest/dbs"
	"github.com/salemzii/tykTest/files"
)

func main() {

	dataLs := files.Reader()

	dbs.WriteData(dataLs)
	fmt.Println("Complete")
}
