package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charly3pins/feedly"
)

const (
	baseURL = "https://sandbox7.feedly.com"
	version = "v3"
)

func main() {
	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		log.Fatal("ACCESS_TOKEN env var missing")
	}
	cli := feedly.Client{
		Config: feedly.ClientConfig{
			BaseURL: baseURL,
			Version: version,
			Token:   token,
		},
		Client: http.Client{
			Timeout: 20 * time.Second,
		},
	}
	resp, err := cli.GetProfile()
	if err != nil {
		log.Fatal(err)
	}
	enc, err := json.MarshalIndent(resp, "  ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(enc))
}
