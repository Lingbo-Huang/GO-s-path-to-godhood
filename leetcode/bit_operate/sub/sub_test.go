package sub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSub(t *testing.T) {
	res := Sub(22, 5)
	assert.Equal(t, 17, res)
}
