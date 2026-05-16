package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func InitLogger() error {
	file, err := os.OpenFile(
		"app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	if err != nil {
		return err
	}

	Log = log.New(file, "INFO: ", log.Ldate|log.Ltime)

	return nil
}
