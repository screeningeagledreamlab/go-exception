package exception

import (
	"testing"

	assert "github.com/screeningeagledreamlab/go-assert"
)

func TestGetStackTrace(t *testing.T) {
	assert := assert.New(t)

	assert.NotEmpty(GetStackTrace())
}
