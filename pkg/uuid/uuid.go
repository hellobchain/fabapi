package uuid

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	return uuid.New().String()
}

func GetUUIDInt() uint32 {
	return uuid.New().ID()
}
