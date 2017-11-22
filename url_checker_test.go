package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURLCheckerCheck(t *testing.T) {
	c := newURLChecker(0, newSemaphore(1024))

	for _, u := range []string{"https://google.com", "README.md"} {
		assert.Equal(t, nil, c.Check(u, "README.md"))
	}

	for _, u := range []string{"https://hey-hey-hi-google.com", "READYOU.md", "://"} {
		assert.NotEqual(t, nil, c.Check(u, "README.md"))
	}
}

func TestURLCheckerCheckWithTimeout(t *testing.T) {
	c := newURLChecker(30*time.Second, newSemaphore(1024))

	for _, u := range []string{"https://google.com", "README.md"} {
		assert.Equal(t, nil, c.Check(u, "README.md"))
	}

	for _, u := range []string{"https://hey-hey-hi-google.com", "READYOU.md", "://"} {
		assert.NotEqual(t, nil, c.Check(u, "README.md"))
	}
}

func TestURLCheckerCheckMany(t *testing.T) {
	c := newURLChecker(0, newSemaphore(1024))

	for _, us := range [][]string{{}, {"https://google.com", "README.md"}} {
		rc := make(chan urlResult, 1024)
		c.CheckMany(us, "README.md", rc)

		for r := range rc {
			assert.NotEqual(t, "", r.url)
			assert.Equal(t, nil, r.err)
		}
	}
}