package internal

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	processList := scanner()
	assert.NotNil(t, processList)
	p := processList[0]
	assert.Equal(t, reflect.TypeOf(p.Id).Kind(), reflect.Int)
	assert.NotEqual(t, p.Name,  "")
}
