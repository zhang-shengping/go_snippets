package main

import (
	"context"
	"flag"
	"fmt"

	// like a util package
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/var/run/kubernetes/admin.kubeconfig", "location to kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", nodes)

	for _, node := range nodes.Items {
		fmt.Printf("%+v\n", node.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, deployment := range deployments.Items {
		fmt.Printf("%+v\n", deployment.Name)
	}
}
