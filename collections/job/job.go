package job

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Job struct {
	Id      uuid.UUID `bson:"id"`
	Status  string    `bson:"status"`
	Command string    `bson:"command"`
	Args    []string  `bson:"args"`

	Response string `bson:"response"`
}

func (j *Job) Encode() ([]byte, error) {
	bsonData, err := bson.Marshal(j)
	if err != nil {
		return nil, err
	}
	return bsonData, nil
}

func (j *Job) Decode(data []byte) error {
	err := bson.Unmarshal(data, j)
	if err != nil {
		return err
	}
	return nil
}
