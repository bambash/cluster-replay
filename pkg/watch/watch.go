package watch

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bambash/cluster-replay/pkg/kube"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
)

func watchPods() {
	clientset := kube.GetClientsetOrDie()
	watcher, err := clientset.CoreV1().Pods("").Watch(metav1.ListOptions{})

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
		fmt.Println(p.Status)
	}
}
