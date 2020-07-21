package helper

import (
	"context"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

// NodeName returns the node name of a given IP address
func NodeName(clientset *kubernetes.Clientset, ipAddress string) (nodeName string, nodeUID types.UID) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.Error("Failed to list nodes")
	}
	for _, n := range nodes.Items {
		nodeName = n.Name
		nodeUID = n.UID
		internalIP := n.Status.Addresses[0].Address
		if ipAddress == internalIP {
			break
		}
	}
	return nodeName, nodeUID
}
