package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
	file  *os.File
)

func Init(logFile string) {
	var err error
	file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Используем буферизированный вывод
	multiWriter := io.MultiWriter(file, os.Stdout)
	
	Info = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Добавьте функцию для закрытия файла
func Close() {
	if file != nil {
		file.Close()
	}
}