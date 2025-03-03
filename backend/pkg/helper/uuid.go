package helper

import "github.com/google/uuid"

func IsValidUUIDv4(u string) bool {
	parsedUUID, err := uuid.Parse(u)
	if err != nil {
		return false
	}
	return parsedUUID.Version() == 4
}
