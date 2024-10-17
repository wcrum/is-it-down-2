package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/wcrum/is-it-down-v2/collections/job"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri        = "mongodb://root:example@localhost:27017/"
	database   = "isitdown"
	collection = "mock_jobs"

	total_fake_jobs = 100
)

func main() {

	mongoOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(
		options.ServerAPI(options.ServerAPIVersion1),
	)
	mg, err := mongo.Connect(context.TODO(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = mg.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	db := mg.Database(database)

	col := db.Collection(collection)

	websites := []string{
		"https://www.google.com",
		"https://www.instagram.com",
	}

	now := time.Now()
	now = now.Add(time.Duration(-100) * time.Minute)

	for _, site := range websites {

		for x := 0; x < total_fake_jobs; x++ {
			job := job.Job{
				Id:          uuid.New(),
				Status:      "complete",
				Command:     "check-latency",
				Args:        []string{site},
				Response:    strconv.Itoa(rand.IntN(100)),
				CompletedAt: now,
				RecievedAt:  now,
			}

			_, err := col.InsertOne(context.TODO(), job)
			if err != nil {
				return
			}

			now = now.Add(time.Duration(1) * time.Minute)
			fmt.Println("inserted")
		}

	}
}
