package exclude

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New().String()
	return id
}

