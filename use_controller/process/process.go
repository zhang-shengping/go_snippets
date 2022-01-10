package process

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitProcess() <-chan struct{} {
	sigs := make(chan os.Signal, 1)
	stop := make(chan struct{})

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		done := <-sigs
		log.Printf("Recivce signal: %s\n", done)
		log.Printf("Shutdown all routine for process")
		close(stop)
	}()

	return stop
}
