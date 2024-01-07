package test

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"io/ioutil"
	"os"
	"testing"

	"github.com/AtahanPoyraz/config"
)

var (
	conf config.Config
)

func init() {
	c, err := config.ReadConfigFromFile("../config.yml")
	if err != nil {
		log.Printf("[ERROR]: Config file read operation failed: %v", err)
		os.Exit(1)
	}
	conf = c
}

//---[ GET REQUEST ]-------------------------------------------------------------------------------------
func Test_GETRequest(t *testing.T) {
	addr := fmt.Sprintf("http://%s:%d/products/get/", conf.Server.Host, conf.Server.Port)

	res, err := http.Get(addr)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("[ERROR]: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	t.Logf("Response: >\n%s", string(body))
}

//---[ POST REQUEST ]------------------------------------------------------------------------------------
func TestPOSTRequest(t *testing.T) {
	addr := fmt.Sprintf("http://%s:%d/products/post/", conf.Server.Host, conf.Server.Port)

	jsonContent := []byte(`
	{
		"name": "example_name",
		"description": "example_descp",
		"price": 19.99,
		"sku":"example_sku"
	}
	`)

	res, err := http.Post(addr, "application/json", bytes.NewBuffer(jsonContent))
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("[ERROR]: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	t.Logf("Response: >\n%s", string(body))
}

//---[ PUT REQUEST ]-------------------------------------------------------------------------------------
func TestPUTRequest(t *testing.T) {
	addr := fmt.Sprintf("http://%s:%d/products/put/3", conf.Server.Host, conf.Server.Port)

	jsonContent := []byte(`
	{
		"name": "example_name2",
		"description": "example_descp2",
		"price": 19.99,
		"sku":"example_sku2"
	}
	`)

	req, err := http.NewRequest("PUT", addr, bytes.NewBuffer(jsonContent))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	t.Logf("Response: >\n%s", string(body))
}

//---[ DELETE REQUEST ]----------------------------------------------------------------------------------
func Test_DELETERequest(t *testing.T) {
	addr := fmt.Sprintf("http://%s:%d/products/delete/3", conf.Server.Host, conf.Server.Port)

	jsonContent := []byte(`
	{
		"name": "example_name2",
		"description": "example_descp2",
		"price": 19.99,
		"sku":"example_sku2"
	}
	`)

	req, err := http.NewRequest("DELETE", addr, bytes.NewBuffer(jsonContent))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("HTTP request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.StatusCode)
	}

	t.Logf("Response: >\n%s", string(body))
}
