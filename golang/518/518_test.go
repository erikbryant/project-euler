package main

import (
	"log"
	"os"
	"testing"
)

// quiet redirects output to null. Usage: defer quiet()()
func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null) // Also redirect log package output
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr) // Restore log output to stderr
	}
}

func TestS(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{7, 0},
		{11, 0},
		{12, 18},
		{100, 1035},
		{1000, 75019},
		{10000, 4225228},
		{100000, 249551109},
		{1000000, 17822459735},
		{2000000, 64710557505},
		{5000000, 356932880607},
		//{10000000, 1316768308545},
		//{100000000, 100315739184392},
	}

	for _, testCase := range testCases {
		answer := S(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
