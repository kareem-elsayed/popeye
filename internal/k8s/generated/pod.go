package generated

import (
	"github.com/derailed/popeye/internal/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pod represents a Kubernetes Pod.
type Pod struct{}

// List all Pods.
func (*Pod) List(c *k8s.Client, ns string) (*v1.PodList, error) {
	var list v1.PodList
	dial, err := c.Dial()
	if err != nil {
		return &list, err
	}

	return dial.CoreV1().Pods(ns).List(metav1.ListOptions{})
}

// Get a Pod.
func (*Pod) Get(c *k8s.Client, name, ns string) (*v1.Pod, error) {
	var res v1.Pod
	dial, err := c.Dial()
	if err != nil {
		return &res, err
	}

	return dial.CoreV1().Pods(ns).Get(name, metav1.GetOptions{})
}