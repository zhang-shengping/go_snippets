package main

import (
	"flag"
	"time"

	// like a util package

	// "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/zhang-shengping/go_snippets/use_controller/controller"
	"github.com/zhang-shengping/go_snippets/use_controller/info"
	"github.com/zhang-shengping/go_snippets/use_controller/process"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
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
	// ---------------------------
	// Test for Clientset

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	/*
		Init process. Create a stopper
	*/
	stop := process.InitProcess()

	/*
		Create a work queue
	*/
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// utils.GetClient()

	/*
		Create a index event informer, it will not resync
	*/
	informer := info.Indexinfor(clientset)

	/*
		Create controller with 2 times retries
	*/
	ctrl := controller.NewController(clientset, queue, informer, 2)

	/*
		Run controller with 1 goroutine, and keep running
	*/
	ctrl.Run(1, stop)

	/*
		Here use NewSharedIndexInformer
	*/

	/*
		Here use NewSharedInformerFactory
	*/
	// informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	// fmt.Println("informerfactory start and wait")
	// go informerfactory.Start(wait.NeverStop)
	// informerfactory.WaitForCacheSync(wait.NeverStop)

	// podinformer := informerfactory.Core().V1().Pods()
	// // get pods from cache.
	// pods := podinformer.Lister().Pods("default")
	// fmt.Println("Pods are: ", pods)

	// informer := informerfactory.Core().V1().Pods().Informer()
	// informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: func(obj interface{}) {
	// 		pod := obj.(*corev1.Pod)
	// 		fmt.Printf("AddFunc is called, Pod is created %+v\n", pod)
	// 	},
	// 	DeleteFunc: func(obj interface{}) {
	// 		pod := obj.(*corev1.Pod)
	// 		fmt.Printf("DeleteFunc is called, Pod is deleted %+v\n", pod)
	// 	},
	// 	UpdateFunc: func(oldObj interface{}, newObj interface{}) {
	// 		oldPod := oldObj.(*corev1.Pod)
	// 		newPod := newObj.(*corev1.Pod)
	// 		fmt.Printf("UpdateFunc is called, Old Pod is %+v, New Pod is %+v\n", oldPod, newPod)
	// 	},
	// })

	// <-wait.NeverStop
}
