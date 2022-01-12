package gb

import (
	"math/rand"
	"time"
)

type GB interface {
	// Validate the data conforms to GB
	Validate(data string) bool
	// Fake return a fake data conforming to GB
	Fake() string
}

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

func RandByte(bytes []byte) byte {
	return bytes[rand.Intn(len(bytes))]
}
