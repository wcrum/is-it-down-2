/*
todo:
- better logging
- main thread starts NATS
	need to ensure controller <> nats


notes:
CID -> Unique ID for workers
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/wcrum/is-it-down-v2/collections/job"
)

type Server struct {
	mu sync.RWMutex

	nc *nats.Conn

	mg *mongo.Client
	db *mongo.Database

	WorkersAvailable bool
}

func (s *Server) RunWebServer() {

	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "asd")
	})
	http.HandleFunc("POST /test", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8080", nil)
}

func (s *Server) SendJobs() {
	for {
		if !s.WorkersAvailable {
			fmt.Println("No workers available... trying again.")
			time.Sleep(10 * time.Second)
			continue
		}

		// temporary, need to collect websites from user / standard list
		websites := []string{
			"https://www.google.com",
			"https://www.youtube.com",
			"https://www.x.com",
			"https://www.instagram.com",
		}

		for _, w := range websites {
			task := job.Job{}
			task.Id = uuid.New()
			task.Command = "check-latency"
			task.Args = []string{w}

			data, err := task.Encode()
			if err != nil {
				return
			}

			s.nc.Publish("jobs", data)
			fmt.Println("sent job", task)
			time.Sleep(10 * time.Second)
		}

	}
}

func (s *Server) CollectJobs() {
	s.nc.Subscribe("jobs.complete", func(msg *nats.Msg) {
		job := job.Job{}
		job.Decode(msg.Data)

		fmt.Println(job)

		col := s.db.Collection("jobs")

		_, err := col.InsertOne(context.TODO(), job)
		fmt.Println(err)
		fmt.Println("ahhh")
		// store bson data?
	})
}

// request workers reply workers
func (s *Server) GetWorkers() {
	workers := []string{}

	sub, err := s.nc.SubscribeSync("controller")
	if err != nil {
		log.Fatal(err)
	}
	s.nc.Flush()

	// Send the request
	for {
		s.nc.PublishRequest("workers.online", "controller", []byte(""))

		// Wait for a single response
		max := 300 * time.Millisecond
		start := time.Now()
		for {
			msg, err := sub.NextMsg(1 * time.Second)
			if err != nil {
				fmt.Println(err)
				break
			}

			if slices.Contains(workers, string(msg.Data)) {
				break
			}

			workers = append(workers, string(msg.Data))

			s.WorkersAvailable = true

			if time.Since(start) > max {
				break
			}
		}

		if s.WorkersAvailable {
			fmt.Printf("Registered workers %v.\n", workers)
		} else {
			fmt.Printf("No workers connected.")
		}
		time.Sleep(10 * time.Second)
	}
	sub.Unsubscribe()
}

const (
	uri        = "mongodb://root:example@localhost:27017/"
	database   = "isitdown"
	collection = "jobs"
)

func main() {
	var err error
	NATS_SERVER := "localhost:4222"
	server := Server{}
	server.WorkersAvailable = false

	// Connect to NATS
	server.nc, err = nats.Connect(NATS_SERVER,
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			log.Printf("client disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			log.Printf("client reconnected")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			log.Printf("client closed")
		}),
		nats.ConnectHandler(func(_ *nats.Conn) {
			fmt.Println("new connection")
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Connet to Mongo
	mongoOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(
		options.ServerAPI(options.ServerAPIVersion1),
	)
	server.mg, err = mongo.Connect(context.TODO(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = server.mg.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	server.db = server.mg.Database(database)

	go server.RunWebServer()
	go server.GetWorkers()
	go server.SendJobs()
	go server.CollectJobs()

	select {}
}
