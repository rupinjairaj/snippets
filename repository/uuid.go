package repository

import (
	"errors"

	"github.com/google/uuid"
)

type UUID interface {
	GenerateUUID() (string, error)
}

type uuidGen struct{}

func NewUUIDGen() UUID {
	return &uuidGen{}
}

func (u *uuidGen) GenerateUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("Error occurred generating UUID using google/uuid")
	}

	return uuid.String(), nil
}
