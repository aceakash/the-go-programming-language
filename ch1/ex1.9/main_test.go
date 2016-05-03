package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestEnsureHttpPrefixDoesNothingIfPrefixAlreadyPresent(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"http://abc.com", "http://abc.com"},
		{"http://www.ok.co.uk", "http://www.ok.co.uk"},
		{"http://localhost", "http://localhost"},
		{"http://192.168.0.1", "http://192.168.0.1"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.out, ensureHTTP(tt.in))
	}
}

func TestEnsureHttpPrefixAddsHttpPrefixIfNotPresent(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"abc.com", "http://abc.com"},
		{"www.ok.co.uk", "http://www.ok.co.uk"},
		{"localhost", "http://localhost"},
		{"192.168.0.1", "http://192.168.0.1"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.out, ensureHTTP(tt.in))
	}
}
