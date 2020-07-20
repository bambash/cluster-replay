package kube

import (
	"os"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// GetClientsetOrDie returns a new Kubernetes Clientset or dies
func GetClientsetOrDie() *kubernetes.Clientset {
	kubeConf, err := config.GetConfig()

	if err != nil {
		logrus.Error(err, "unable to get Kubernetes client config")
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(kubeConf)

	if err != nil {
		logrus.Error(err, "unable to get Kubernetes clientset")
		os.Exit(1)
	}

	return clientset
}
