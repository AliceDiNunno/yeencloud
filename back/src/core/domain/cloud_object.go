package domain

import "github.com/google/uuid"

type CloudObject struct {
	ID   string
	Name string
}

func (object *CloudObject) Initialize() {
	if object.ID == "" {
		object.ID = uuid.New().String()
	}
}
