package controller

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stop
// channel is closed, at which point it will shutdown the workqueue and wait
// for workers to finish processing their current work items.
// type Controller interface {
// 	Run(stop <-chan struct{})
// 	runWorker()
// 	processNextItem() bool
// 	processItem(key string) error
// }

type Controller struct {
	clientset  kubernetes.Interface
	queue      workqueue.RateLimitingInterface
	informer   cache.SharedInformer
	maxRetries int
}

// We need a constructor function (not method), because all attributes of Controller are private
func NewController(
	clientset kubernetes.Interface,
	queue workqueue.RateLimitingInterface,
	informer cache.SharedInformer,
	maxRetries int,
) *Controller {
	controller := &Controller{
		clientset:  clientset,
		queue:      queue,
		informer:   informer,
		maxRetries: maxRetries,
	}

	log.Printf("Setting up event handlers\n")
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// we get a object here inface, althoght the informer is a sharedindexinformer.
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Printf("AddFunc get object: %v\nEnqueue the key: %s\n",
				obj, key)
			if err == nil {
				// here we put the key into the queue rather then object itself.
				queue.AddRateLimited(key)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			// TODO: reflect the object type ?
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			log.Printf("DeleteFunc get object %v\nDequeue the key: %s\n",
				obj, key)
			if err == nil {
				queue.AddRateLimited(key)
			}
		},
	})
	return controller
}

func (c *Controller) Run(routines int, stop <-chan struct{}) {
	// don't let panics crash the process
	defer runtime.HandleCrash()
	// make sure the work queue is shutdown which will trigger workers to end
	defer c.queue.ShutDown()

	log.Println("Starting test controller")

	// start informer goroutines
	go c.informer.Run(stop)

	// wiat for the caches to synchronize before starting the workers (goroutines ?)
	if !cache.WaitForCacheSync(stop, c.informer.HasSynced) {
		log.Panicln("Time out waiting for caches to sync")
		return
	}

	log.Println("Informer cache synced")

	// run multiple workers
	for i := 0; i < routines; i++ {
		go wait.Until(c.runWorker, time.Second, stop)
	}

	log.Printf("Starting %d workers to consumes the work queue\n", routines)

	// keep controller running in main
	log.Println("Started workers")
	<-stop
	log.Println("Shutting down workers")
}

func (c *Controller) runWorker() {
	// processItem return True continue
	for c.processNextItem() {
		// continue looping
	}
}

// Process with queue
// processNextWorkItem deals with one key off the queue.  It returns false
// when it's time to quit.
func (c *Controller) processNextItem() bool {
	// pull the next work item from queue.  It should be a key we use to lookup
	// something in a cache. (we put the key in informer event handler)
	// if quit is True, that mean the queue shutdown
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// you always have to indicate to the queue that you've completed a piece of
	// work
	defer c.queue.Done(key)

	// do working
	// key is a interface type, convert it to string
	err := c.processItem(key.(string))

	if err == nil {
		log.Printf("Work of %s is done\n", key)
		// No error, tell the queue to stop tracking history
		// Forget indicates that an item is finished being retried.  Doesn't matter whether it's for perm failing
		// or for success, we'll stop the rate limiter from tracking it.  This only clears the `rateLimiter`, you
		// still have to call `Done` on the queue.
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < c.maxRetries {
		log.Printf("Error processing %s (will retry): %v\n", key, err)
		// requeue the item to work on later
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		log.Panicf("Error processing %s (giving up): %v\n", key, err)
		c.queue.Forget(key)
		runtime.HandleError(err)
	}
	return true
}

// Process with objects
func (c *Controller) processItem(key string) error {
	log.Printf("Processing change to resouce %s", key)

	// just for test to convert the informer to different type
	obj, exists, err := c.informer.(cache.SharedIndexInformer).GetIndexer().GetByKey(key)
	if err != nil {
		log.Panicf("Error fetching object with key %s from store: %v\n", key, err)
	}

	if !exists {
		//delete event
		return nil
	}

	log.Printf("Processing object: %+v", obj)
	// !!! create event / update event,
	// when update we compare the ResourceVersion with UpdateFunc
	// but if use the NewSharedIndexInformer, NewSharedInformerFactory?
	// can we still get object.ResourceVersion?
	return nil
}
