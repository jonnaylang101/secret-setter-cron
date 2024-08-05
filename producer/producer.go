package main

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/sirupsen/logrus"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Set a k8s secret periodically
	logger := logrus.New()
	logger.Info("+++ Producer started +++")

}

func setSecret(logger *logrus.Logger) {
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

	username := fmt.Sprintf("BarnabyClunkwurst-%d", rand.IntN(1000))

	secret := &coreV1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: "producer-secret",
		},
		Data: map[string][]byte{
			"username": []byte(username),
		},
	}

	_, err = clientset.CoreV1().Secrets("default").Update(context.Background(), secret, v1.UpdateOptions{})
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
