package main

import (
	"learngo/self_interview/study_drawdraw/study0617/log_go_web/logging"
)

func writeMessage(logger logging.Logger) {
	logger.Info("Let's Go, logger")
}

func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}
