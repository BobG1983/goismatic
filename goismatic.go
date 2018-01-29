package goismatic

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type lang int

// The list of supported languages
const (
	English lang = iota
	Russian
)

var languages = [...]string{
	"en",
	"ru",
}

// The Quote struct represents the details of a quote from forismatic.com
type Quote struct {
	QuoteText   string
	QuoteAuthor string
}

func (q *Quote) String() string {
	return q.QuoteText + " - " + q.QuoteAuthor
}

const apiURL string = "http://api.forismatic.com/api/"
const apiVer string = "1.0"

// Get returns a random quote in either English or Russian from forismatic.com
func Get(l lang) (*Quote, error) {
	q := new(Quote)

	URL := apiURL + apiVer + "/"

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	qu := req.URL.Query()
	qu.Add("method", "getQuote")
	qu.Add("format", "json")
	qu.Add("lang", languages[l])
	req.URL.RawQuery = qu.Encode()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(q)
	if err != nil {
		return nil, err
	}

	if q.QuoteAuthor == "" {
		q.QuoteAuthor = "Unknown"
	}

	return q, nil
}
