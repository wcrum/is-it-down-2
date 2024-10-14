// Listen for new connections
// Listen for disconnections
//     this was handled by the request reply

// Send jobs to available connections

// Future -> start NATS Server

// CID -> Unique ID for workers

package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/wcrum/is-it-down-v2/collections/job"
)

type Server struct {
	mu sync.RWMutex

	nc *nats.Conn

	WorkersAvailable bool
}

func (s *Server) RunWebServer() {
	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "asd")
	})
	http.HandleFunc("POST /test", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8080", nil)
}

/*
google.com
	| status
	| latency
	| last_checked
	| history {
		- time:
		- time:
		- time:
		- time:
		- time:
	}
*/

func (s *Server) SendJobs() {
	for {
		if !s.WorkersAvailable {
			fmt.Println("No workers available... trying again.")
			time.Sleep(10 * time.Second)
			continue
		}
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

const uri = "mongodb://root:example@localhost:27017/"

func main() {
	var err error
	NATS_SERVER := "localhost:4222"
	server := Server{}
	server.WorkersAvailable = false

	go server.RunWebServer()

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

	go server.GetWorkers()
	go server.SendJobs()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	/*
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
		// Send a ping to confirm a successful connection

		db := client.Database("test")

		type Test struct {
			CreatedAt time.Time `bson:"time"`
		}
		test := Test{CreatedAt: time.Now().UTC()}

		res, err := db.Collection("test").InsertOne(context.Background(), test)
		if err != nil {
			log.Fatal(err)
		}

		bson.Marshal()
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
		log.Printf("%v documents inserted", res.InsertedID)
	*/

	select {}
}
