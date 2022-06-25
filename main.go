package main

import (
	"fmt"

	"github.com/salemzii/tykTest/dbs"
	"github.com/salemzii/tykTest/files"
)

var (
//waitgroup sync.WaitGroup
)

func main() {

	dataLs := files.Reader()
	fmt.Println(dataLs)
	dbs.WriteData(dataLs)
	fmt.Println("Complete")
}
