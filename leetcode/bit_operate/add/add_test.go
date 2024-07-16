package add

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(5, 17)
	assert.Equal(t, 22, res)
}
