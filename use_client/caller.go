package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	// like a util package
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/var/run/kubernetes/admin.kubeconfig", "location to kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	config.Timeout = 120 * time.Second
	if err != nil {
		// handle error, if can not find the kubeconfig
		config, err = rest.InClusterConfig()
		if err != nil {
			panic("can not find serviceaccount config")
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// k8s API object implements runtime.Object
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", nodes)

	// for _, node := range nodes.Items {
	// 	fmt.Printf("%+v\n", node.Name)
	// }

	// deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, deployment := range deployments.Items {
	// 	fmt.Printf("%+v\n", deployment.Name)
	// }

	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	fmt.Println("informerfactory start and wait")
	go informerfactory.Start(wait.NeverStop)
	informerfactory.WaitForCacheSync(wait.NeverStop)

	podinformer := informerfactory.Core().V1().Pods()
	// get pods from cache.
	pods := podinformer.Lister().Pods("default")
	fmt.Println("Pods are: ", pods)

	informer := informerfactory.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(corev1.Pod)
			fmt.Printf("AddFunc is called, Pod is created %+v\n", pod)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(corev1.Pod)
			fmt.Printf("DeleteFunc is called, Pod is deleted %+v\n", pod)
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			oldPod := oldObj.(corev1.Pod)
			newPod := newObj.(corev1.Pod)
			fmt.Printf("UpdateFunc is called, Old Pod is %+v, New Pod is %+v\n", oldPod, newPod)
		},
	})

	// Create a channel to stops the shared informer gracefully
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()

	go informer.Run(stopper)

	// podinformer.Informer().AddEventHandler(
	// 	cache.ResourceEventHandlerFuncs{

	// 	}
	// )

	// You need to start the informer, in my case, it runs in the background
}
