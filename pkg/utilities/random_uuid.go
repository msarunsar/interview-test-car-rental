package utilities

import "github.com/google/uuid"

func GenerateRandomUUID() string {
	return uuid.New().String()
}

type UUIDGenerator interface {
	GenerateRandomUUID() string
}

type DefaultUUIDGenerator struct{}

func (d DefaultUUIDGenerator) GenerateRandomUUID() string {
	return uuid.New().String()
}

var UUIDGeneratorInstance UUIDGenerator = DefaultUUIDGenerator{}
