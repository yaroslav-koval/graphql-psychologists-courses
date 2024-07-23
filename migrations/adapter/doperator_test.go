package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExist(t *testing.T) {
	do := NewDirectoryOperator()
	ok := do.IsFileExist("this file is not exists")
	assert.False(t, ok)
}
