package watch

import (
	"github.com/bambash/cluster-replay/pkg/kube"
)

func WatchResources() {
	clientset := kube.GetClientsetOrDie()
	watchPods(clientset)
	watchNodes(clientset)
}
