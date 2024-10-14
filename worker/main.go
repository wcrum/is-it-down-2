// Listen for new connections
// Listen for disconnections

// Send jobs to available connections

// Future -> start NATS Server

// CID -> Unique ID for workers

/*
Server SUBJ -> worker.<cid>.job
Server SUBJ -> worker.<cid>.job.<jid>
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/wcrum/is-it-down-v2/collections/job"
)

var (
	InvalidJob = fmt.Errorf("Invalid Job.")
)

type Worker struct {
	mu sync.RWMutex

	nc *nats.Conn
}

func (w *Worker) CheckLatency(args []string) (time.Duration, error) {
	fmt.Println("")
	if !(len(args) >= 1) {
		return 0, InvalidJob
	}

	start := time.Now()

	url := args[0]

	// task
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	latency := time.Since(start)

	fmt.Println(latency)

	return latency, nil
}

func (w *Worker) CheckOnline() {
	sub, err := w.nc.SubscribeSync("workers.online")
	if err != nil {

	}

	for {
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
		}

		me, err := w.nc.GetClientID()
		if err != nil {
		}

		clientID := fmt.Sprintf("%d", me)

		msg.Respond([]byte(clientID))
	}
}

func main() {
	var err error

	NATS_SERVER := "localhost:4222"

	worker := Worker{}
	worker.nc, err = nats.Connect(NATS_SERVER)
	if err != nil {
		return
	}

	go worker.CheckOnline()

	go func() {
		worker.nc.QueueSubscribe("jobs", "workers", func(msg *nats.Msg) {
			task := job.Job{}

			if err := task.Decode(msg.Data); err != nil {
				return
			}

			fmt.Println(task)
			// handle job
			switch task.Command {
			case "check-latency":
				go worker.CheckLatency(task.Args)
			}

		})
	}()

	select {}
}
