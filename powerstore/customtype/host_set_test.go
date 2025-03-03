package customtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostSetNormaization(t *testing.T) {
	out, err := NewHostSetType().normalizeStrings([]string{"192.168.1.0/255.255.255.0"})
	assert.NoError(t, err)
	assert.Equal(t, []string{"192.168.1.0/24"}, out)
}
