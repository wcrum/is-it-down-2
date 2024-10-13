package job

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Job struct {
	Id      uuid.UUID `json:"id"`
	Command string    `json:"command"`
	Args    []string  `json:"args"`
}

func (j *Job) Bytes() ([]byte, error) {
	jsonData, err := json.Marshal(j)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}
