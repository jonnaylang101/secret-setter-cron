package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Set a k8s secret periodically
	logger := logrus.New()
	logger.Info("+++ Producer started +++")
}
