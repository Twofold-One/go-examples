package basic

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func LoggingExmaple() {
	f, err := os.OpenFile("/home/twofold_one/GitProjects/go/go-examples/basic/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("System is starting")
	// log.Panicln("System stop")

	// Logs using logrus module
	logrus.Println("Logrus output")

	// Levels
	logrus.Info("Info message")
	logrus.Warn("Warning!")
	logrus.Error("Some error")

	// Fields
	workerLogger := logrus.WithFields(logrus.Fields{
		"source": "worker",
	})
	workerLogger.Info("worker has finished processed task")

}
