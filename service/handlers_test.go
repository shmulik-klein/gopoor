package service

import (
	"bytes"
	"fmt"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateBuy(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(createBuyHandler(formatter)))
	defer server.Close()

	body := []byte("{\n \"name\": \"grocery\"\n}")
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error in creating POST request for createBuyHandler: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createBuyHandler: %v", err)
	}
	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received: %s", res.Status)
	}

	fmt.Printf("Payload: %s", payload)
}
