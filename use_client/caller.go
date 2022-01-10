package main

import (
	"flag"
	"fmt"
	"time"

	// like a util package

	// "k8s.io/apimachinery/pkg/util/runtime"
	corev1 "k8s.io/api/core/v1"
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
	// ---------------------------
	// Test for Clientset

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// k8s API object implements runtime.Object
	// nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", nodes)

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

	// -------------------
	// Test for Clientset get informerfactory and event

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
			pod := obj.(*corev1.Pod)
			fmt.Printf("AddFunc is called, Pod is created %+v\n", pod)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			fmt.Printf("DeleteFunc is called, Pod is deleted %+v\n", pod)
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			oldPod := oldObj.(*corev1.Pod)
			newPod := newObj.(*corev1.Pod)
			fmt.Printf("UpdateFunc is called, Old Pod is %+v, New Pod is %+v\n", oldPod, newPod)
		},
	})

	// --------------
	// rootCAs, _ := x509.SystemCertPool()
	// certs, _ := ioutil.ReadFile("/var/run/kubernetes/server-ca.crt")
	// rootCAs.AppendCertsFromPEM(certs)
	// tlsconfig := &tls.Config{
	// 	// InsecureSkipVerify: *insecure,
	// 	RootCAs: rootCAs,
	// }

	// tr := http.Transport{TLSClientConfig: tlsconfig}

	// Test for Clientset get informer, cache
	// config.GroupVersion = &metav1.SchemeGroupVersion
	// config.APIPath = "api"
	// config.GroupVersion = &corev1.SchemeGroupVersion
	// config.NegotiatedSerializer = scheme.Codecs
	// config.CertFile = "/var/run/kubernetes/client-admin.crt"
	// config.KeyFile = "/var/run/kubernetes/client-admin.key"
	// config.CAFile = "/var/run/kubernetes/server-ca.crt"
	// // config.NegotiatedSerializer = runtime.NewSimpleNegotiatedSerializer()
	// // fmt.Printf("CAFILE is %+v \n", config.CAFile)
	// config.TLSClientConfig.CAFile
	// restclient, err := rest.RESTClientForConfigAndClient(config, nil)
	// restclient.Client.Transport = tr
	// restclient, err := rest.UnversionedRESTClientFor(config)

	/*
		I0107 06:17:09.400481   24798 httplog.go:129] "HTTP" verb="LIST" URI="/api/v1/namespaces/default/pods?timeout=2m0s" latency="2.168844ms" userAgent="Go-http-client/2.0" audit-ID="7f76d6dc-8c9c-4c97-ad79-114be6144244" srcIP="[::1]:39696" resp=200
	*/

	// restclient, err := rest.HTTPClientFor(config)

	// restclient, err := rest.RESTClientFor(config)
	// restclient, err := clientv1.NewForConfig(config)

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("can not build restclient")
	// }

	// lswch := cache.NewListWatchFromClient(
	// 	restclient,
	// 	"pods",
	// 	"default",
	// 	// fields.Nothing(),
	// 	fields.Everything(),
	// )

	// obj, err := lswch.List(metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("list error")
	// }
	// fmt.Println(obj)
	// lswch.Watch(metav1.ListOptions{})

	//----------------
	<-wait.NeverStop
}
