package main

import (
	"fmt"
	"sync"

	"github.com/salemzii/tykTest/dbs"
	"github.com/salemzii/tykTest/files"
)

var (
	waitgroup sync.WaitGroup
)

func main() {

	dataLs := files.Reader()
	waitgroup.Add(len(dataLs))

	for _, v := range dataLs {
		go func() {
			defer waitgroup.Done()
			dbs.WriteData(&v)
		}()
	}
	waitgroup.Wait()
	fmt.Println("Complete")
}
