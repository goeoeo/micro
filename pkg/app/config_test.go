package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := Config{
		Mode: ModeDev,
		Host: "127.0.0.1",
		Port: "8888",
	}

	assert.True(t, c.IsDev())

	assert.Equal(t, c.Endpoint(), "127.0.0.1:8888")
}
