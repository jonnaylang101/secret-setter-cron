package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Get K8s sercret and log it
	logger := logrus.New()
	logger.Info("+++ Consumer started +++")
}
