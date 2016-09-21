package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	os.Clearenv()
	os.Setenv("STATSD_HOST", "1.2.3.4")
	os.Setenv("STATSD_PORT", "1234")
	os.Setenv("STATSD_PREFIX", "unit-test")

	config := getConfig()

	assert.Equal(t, "1.2.3.4", config.Host, "should read STATSD_HOST")
	assert.Equal(t, 1234, config.Port, "should read STATSD_PORT")
	assert.Equal(t, "unit-test", config.Prefix, "should read STATSD_PREFIX")
}
