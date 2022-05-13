package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet(t *testing.T) {
	got := NeedsAComment()

	assert.Equal(t, "Hi!", got, "should properly greet")
}
