package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// BenchmarkRESTLookup runs a benchmark test of the REST GET endpoint
func BenchmarkRESTLookup(b *testing.B) {
	client := &http.Client{}

	// run http posts against it
	for i := 0; i < b.N; i++ {
		_, err := client.Get("http://localhost:8889/lookup/1")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRESTSet runs a benchmark test of the REST POST endpoint
func BenchmarkRESTSet(b *testing.B) {
	client := &http.Client{}

	// run http posts against it
	for i := 0; i < b.N; i++ {
		req, err := json.Marshal(&SetRequest{Entity: &Entity{Id: 1234, Name: "Small Entity"}})
		if err != nil {
			b.Fatal(err)
		}

		_, err = client.Post("http://localhost:8889/set", "application/json", bytes.NewReader(req))
		if err != nil {
			b.Fatal(err)
		}
	}
}
