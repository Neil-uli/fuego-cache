package cache

import (
	"bytes"
	"encoding/gob"
	"github.com/tomiok/fuego-cache/logs"
	"math"
)

const (
	prime = 127
	m     = math.MaxInt64
)

func ApplyHash(i interface{}) int {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)

	if err != nil {
		logs.Error("cannot encode interface: " + err.Error())
		return 0
	}
	byteValues := buf.Bytes()
	var hashcode int
	for i, v := range byteValues {
		hashcode += prime*i + int(v)%m
	}

	return hashcode
}
