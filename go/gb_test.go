package gb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGB(t *testing.T) {
	for name, gb := range gbs {
		n := random.Intn(1000)
		for i := 0; i < n; i++ {
			data := gb.Fake()
			assert.True(t, gb.Validate(data), "err from %v, data is %v", name, data)
		}
	}
}
