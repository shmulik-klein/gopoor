package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreatePurchaseReturnsStatusOK(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(purchaseHandler(formatter))
	defer server.Close()

	body := []byte("{\n \"name\": \"supermarket\",\n \"price\": \"1000\"\n}")
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createPurchaseHandler: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createPurchaseHandler: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response code 201, received: %v", res.StatusCode)
	}
	header := res.Header
	if header["Location"] == nil {
		t.Errorf("Expected response to include Location header")
	}
}
