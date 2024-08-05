package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Get K8s sercret and log it
	logger := logrus.New()
	logger.Info("+++ Consumer started +++")
	ticker := time.NewTicker(20 * time.Second)
	count := 0

	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Error("Error building kubeconfig: ", err)
		return
	}

	logger.Infof("Kubeconfig: %+v", config)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error("Error building clientset: ", err)
		return
	}

	for range ticker.C {
		count++
		logger.Infof("+++ Consumer tick %d +++", count)

		secret, err := clientset.CoreV1().Secrets("default").Get(context.Background(), "producer-secret", v1.GetOptions{})
		if err != nil {
			logger.Error("Error getting secret: ", err)
			return
		}

		byt, ok := secret.Data["username"]
		if !ok {
			logger.Error("Error getting username from secret")
			return
		}

		logger.Info("Username: ", string(byt))
	}
}
