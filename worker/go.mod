module github.com/wcrum/is-it-down-2/worker

go 1.21.9

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/nats-io/nats.go v1.37.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/wcrum/is-it-down-v2/collections/job v0.0.0-00010101000000-000000000000 // indirect
	go.mongodb.org/mongo-driver v1.17.1 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
)

replace github.com/wcrum/is-it-down-v2/collections/job => ../collections/job
