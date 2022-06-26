package logger

import "log"

func ErrorLogger(errs error) {
	file, err = openLogFile("./myerror.log")
	if err != nil {
		log.Fatal(err)
	}
	errorLog := log.New(file, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Println(errs)
}
