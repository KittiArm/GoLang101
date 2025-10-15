package uuid

import (
	"github.com/google/uuid"
)

func genUUID() string {
	userUUID := uuid.New()

	return userUUID.String()
}