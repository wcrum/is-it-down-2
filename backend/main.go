package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wcrum/watcher/collections/job"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func listenToRunners(sub *nats.Subscription, server *Server) {
	for {
		msg, err := sub.NextMsg(time.Second)
		if err != nil {
			if err == nats.ErrTimeout {
				// No message received within the timeout period
				continue
			}
			log.Fatal(err)
		}

		if msg.Subject == "worker.disconnect" {
			for i, w := range server.Workers {
				id, err := uuid.Parse(string(msg.Data))
				if err != nil {
				}
				if w.ID == id {
					server.Workers = append(server.Workers[:i], server.Workers[i+1:]...)
					break
				}
			}
		}

		if msg.Subject == "worker.connect" {
			server.mu.Lock()
			id, err := uuid.Parse(string(msg.Data))
			if err != nil {
			}
			worker := Worker{
				ID: id,
			}
			server.Workers = append(server.Workers, worker)
			server.mu.Unlock()
		}

		// Print the received message.
		fmt.Printf("%s: %s\n", string(msg.Subject), string(msg.Data))
	}
}

type Server struct {
	mu      sync.RWMutex
	NATS    *nats.Conn
	Workers []Worker
	Jobs    Jobs
}

type Worker struct {
	ID   uuid.UUID
	Busy bool
}

type Jobs struct {
	Completed []job.Job
	Sent      []job.Job
	Queued    []job.Job
}

func (j *Jobs) getJob() job.Job {
	job := j.Queued[0]

	j.Queued = j.Queued[1:]

	return job
}

func (s *Server) handleJobs() {

	w := 0
	for {
		wait := false

		if len(s.Workers) == 0 {
			fmt.Println("No workers found. Waiting 5 seconds.")
			wait = true
		}

		if len(s.Jobs.Queued) == 0 {
			fmt.Println("No jobs found. Waiting 5 seconds.")
			wait = true
		}

		if wait {
			time.Sleep(5 * time.Second)
			continue
		}

		if w >= len(s.Workers) {
			w = 0
		}

		worker := s.Workers[w]
		s.sendJob(worker, s.Jobs.getJob())

		w++
		time.Sleep(2 * time.Second)

	}
}

func (s *Server) sendJob(worker Worker, job job.Job) {
	subj := fmt.Sprintf("worker.%s.job", worker.ID)
	fmt.Printf("Sending job: %s\n", subj)

	b, err := job.Bytes()
	if err != nil {
	}
	s.NATS.Publish(subj, b)
}

func (s *Server) newJob() {
	job := job.Job{
		Id:      uuid.New(),
		Command: "ping",
		Args:    []string{"google.com"},
	}
	s.Jobs.Queued = append(s.Jobs.Queued, job)
}

func (s *Server) mockJobs() {
	for {
		s.newJob()
		fmt.Println(len(s.Jobs.Queued))
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var err error

	server := Server{
		Workers: []Worker{},
		Jobs: Jobs{
			Completed: []job.Job{},
			Sent:      []job.Job{},
			Queued:    []job.Job{},
		},
	}

	server.NATS, err = nats.Connect("localhost:4222", nats.Name("API PublishBytes Example"))

	defer server.NATS.Close()

	if err := server.NATS.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}

	sub, err := server.NATS.SubscribeSync("worker.*")
	if err != nil {
		log.Fatal(err)
	}

	go listenToRunners(sub, &server)

	go server.handleJobs()

	go server.mockJobs()

	select {}
}
