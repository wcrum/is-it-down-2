package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/wcrum/watcher/collections/job"
)

func main() {
	id := uuid.New()
	name := id.String()

	fmt.Printf("Starting worker %s\n", name)

	var wg sync.WaitGroup
	wg.Add(1) // Increment WaitGroup counter
	// Connect to NATS server.
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Publish("worker.connect", []byte(name))

	// Subscribe to a subject.

	jobSubscription := fmt.Sprintf("worker.%s.job", name)
	fmt.Println(jobSubscription)
	sub, err := nc.SubscribeSync(jobSubscription)
	if err != nil {
		log.Fatal(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT)
	go func() {
		defer wg.Done() // Decrement counter when goroutine exits
		s := <-sigchan
		fmt.Println(s)
		err := nc.Publish("worker.disconnect", []byte(name))
		if err != nil {
			fmt.Println("Error publishing message:", err)
		}
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}()

	for {
		msg, err := sub.NextMsg(time.Second)
		if err != nil {
			if err == nats.ErrTimeout {
				// No message received within the timeout period
				continue
			}
			log.Fatal(err)
		}

		// Print the received message.
		go func() {
			job := job.Job{}
			err := json.Unmarshal(msg.Data, &job)

			fmt.Println(err)
			fmt.Printf("Received message: %s\n", job)
			time.Sleep(1 * time.Second)
			fmt.Println("Finished workload.")
		}()

		// Process the message...
	}
}
