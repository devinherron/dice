package dice

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
)

func Roll(size int) int {
	var src cryptoSource
	rnd := rand.New(src)

	return rnd.Intn(size) + 1
}

func RollMulti(size int, number int) ([]int, int) {
	var sum int
	results := []int{}

	var src cryptoSource
	rnd := rand.New(src)

	for i := 0; i < number; i++ {
		results = append(results, rnd.Intn(size)+1)
		sum += results[i]
	}
	return results, sum
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
