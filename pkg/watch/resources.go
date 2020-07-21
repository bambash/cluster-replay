package watch

import (
	"context"
	"fmt"

	"github.com/bambash/cluster-replay/pkg/helper"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
)

func watchPods(clientset *kubernetes.Clientset) {
	watcher, err := clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})

	if err != nil {
		logrus.Error(err, "Unable to watch for Pods")
		runtime.HandleError(err)
	}

	ch := watcher.ResultChan()

	for event := range ch {
		p, ok := event.Object.(*v1.Pod)
		if !ok {
			logrus.Error("Could not parse Pod object")
		}
		nodeName, _ := helper.NodeName(clientset, p.Status.HostIP)
		fmt.Println(p.Name, p.Namespace, p.UID, nodeName, p.Labels, p.Status.Phase)
	}
}

func watchNodes(clientset *kubernetes.Clientset) {
	watcher, err := clientset.CoreV1().Nodes().Watch(context.TODO(), metav1.ListOptions{})

	if err != nil {
		logrus.Error(err, "Unable to watch for Nodes")
		runtime.HandleError(err)
	}

	ch := watcher.ResultChan()

	for event := range ch {
		n, ok := event.Object.(*v1.Node)
		if !ok {
			logrus.Error("Could not parse Node object")
		}
		fmt.Println(n.Name, n.UID, n.Status.Phase)
	}
}
