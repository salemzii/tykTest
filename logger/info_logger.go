package logger

import "log"

// Logs successful calls to myinfo.log
func InfoLogger(info string) {
	file, err = openLogFile("./myinfo.log")
	if err != nil {
		log.Fatal(err)
	}
	infoLog := log.New(file, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Println(info)
}
