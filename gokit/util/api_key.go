package util

import "github.com/google/uuid"

func NewKeyString() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uuid_str := uuid.String()
	return uuid_str, nil
}
