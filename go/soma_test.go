package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_soma(t *testing.T) {
	t.Run("OK", func(t *testing.T){
		result, err := soma("1", "2")
		assert.NotNil(t, err)
		assert.Equals(t, 3., result)
	})
}
