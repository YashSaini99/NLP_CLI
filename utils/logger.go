package utils

import (
	"log"
)

// InitializeLogger sets up logging with a standard format.
func InitializeLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Uncomment the following lines to log to a file instead of stderr.
	/*
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	*/
}
