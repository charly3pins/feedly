package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	redirectURL  = "http://localhost:8080/"
	authorizeURL = "https://sandbox7.feedly.com/v3/auth/auth"
	tokenURL     = "https://sandbox7.feedly.com/v3/auth/token"
)

type TokenHandler struct {
	config *oauth2.Config
}

func NewTokenHandler(clientID, clientSecret, redirectURL string, scopes ...string) TokenHandler {
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeURL,
			TokenURL: tokenURL,
		},
	}

	return TokenHandler{
		config: &config,
	}
}

func (t TokenHandler) Token(state string, r *http.Request) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, fmt.Errorf("error auth failed: %+v", e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("error no access code present")
	}
	urlState := values.Get("state")
	if urlState != state {
		return nil, errors.New("error state parameter doesn't match")
	}
	return t.config.Exchange(context.Background(), code)
}

func (t TokenHandler) AuthURL(state string) string {
	return t.config.AuthCodeURL(state)
}

var (
	th    = NewTokenHandler("sandbox", "Vl8UzizSQzmlorcBqEhT9VX32Gn5jzTt", redirectURL, "https://cloud.feedly.com/subscriptions")
	ch    = make(chan *oauth2.Token)
	state = "superSecret"
)

func main() {
	http.HandleFunc("/", completeAuth)
	go http.ListenAndServe(":8080", nil)

	url := th.AuthURL(state)
	fmt.Println("Log in to Feedly visiting:", url)

	t := <-ch
	fmt.Println(t)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	token, err := th.Token(state, r)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "error getting token", http.StatusForbidden)
	}
	if st := r.FormValue("state"); st != state {
		log.Fatalf("error state expected: %s - state recieved %s\n", state, st)
		http.NotFound(w, r)
	}
	ch <- token
}
