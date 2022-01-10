package info

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

/*
Pod index informer for event, so we skip resync here. to set resync 0
we still get a object when event / resync, althoght the informer is a sharedindexinformer.
but we can use index to trace object in store, it is effiecient when the object is large
*/
func Indexinfor(clientset kubernetes.Interface) cache.SharedInformer {
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				// return clientset.CoreV1().Pods(metav1.NamespaceAll).List(options)
				return clientset.CoreV1().Pods(metav1.NamespaceAll).List(
					context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return clientset.CoreV1().Pods(metav1.NamespaceAll).Watch(
					context.Background(), options)
			},
		},
		&corev1.Pod{},
		0, //Skip resync
		cache.Indexers{},
	)
	return informer
}
