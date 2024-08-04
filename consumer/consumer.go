package main

import (
	"context"

	"github.com/sirupsen/logrus"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Get K8s sercret and log it
	logger := logrus.New()
	logger.Info("+++ Consumer started +++")

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

	secret, err := clientset.CoreV1().Secrets("default").Get(context.Background(), "producer-secret-8d7chtfcmb", v1.GetOptions{})
	if err != nil {
		logger.Error("Error getting secret: ", err)
		return
	}

	logger.Infof("Secret: %+v", secret)

	byt, ok := secret.Data["username"]
	if !ok {
		logger.Error("Error getting username from secret")
		return
	}

	logger.Info("Username: ", string(byt))
}
