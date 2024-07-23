package logging

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMarshalZerologArray(t *testing.T) {
	res := new(logResult)
	zl := zerolog.New(res)
	SetLogger(zl)

	st := StackTrace{"error1", "error2"}
	zl.Info().Array("stack", &st).Send()

	v, ok := res.data["stack"]
	assert.True(t, ok)
	vArray, ok := v.([]interface{})
	assert.True(t, ok)
	assert.Equal(t, 2, len(vArray))
	assert.Equal(t, st[0], vArray[0])
	assert.Equal(t, st[1], vArray[1])
}
