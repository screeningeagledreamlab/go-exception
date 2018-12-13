package exception

import (
	"testing"

	assert "github.com/apremalal/go-assert"
)

func TestGetStackTrace(t *testing.T) {
	assert := assert.New(t)

	assert.NotEmpty(GetStackTrace())
}
