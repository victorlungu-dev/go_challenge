package http

import (
	"azarc.io/internal/imdb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
)

// MaxRequestsReachedError is the error returned by the http repo when the requests made to the api reach the limit
var MaxRequestsLimitError error = maxRequestsLimitError{}

type maxRequestsLimitError struct{}

func (maxRequestsLimitError) Error() string { return "max requests limit reached" }

type HttpRepo struct {
	ctx          context.Context
	url          *url.URL
	client       *http.Client
	apiKey       string
	requestCount int32
	MaxRequests  int
}

func (hr *HttpRepo) Retrieve(id string) (*imdb.Movie, error) {
	if hr.MaxRequestsReached() {
		return nil, maxRequestsLimitError{}
	}
	urlCopy := *hr.url
	q := urlCopy.Query()
	q.Set("apiKey", hr.apiKey)
	q.Set("i", id)
	urlCopy.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, urlCopy.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/json")
	req = req.WithContext(hr.ctx)
	res, err := hr.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, err
		}

		return nil, err

	}
	hr.incrementReqCounter()
	defer res.Body.Close()

	if res.StatusCode == 403 {
		return nil, errors.New(fmt.Sprintf("invalid api_key or rate limit reached you should pass maxRequests to the command"))
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading body: %v", err))
	}

	var m imdb.Movie
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error unmarshalling data for %s: %v", id, err))
	}
	return &m, nil
}

func (hr *HttpRepo) incrementReqCounter() {
	atomic.AddInt32(&hr.requestCount, 1)
}

func (hr *HttpRepo) MaxRequestsReached() bool {
	return hr.MaxRequests <= int(hr.requestCount)
}

func NewHttpRepo(ctx context.Context, baseUrl, apiKey string, maxRequests int) imdb.MovieRepo {
	u, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	return &HttpRepo{ctx: ctx, url: u, client: client, apiKey: apiKey, MaxRequests: maxRequests}
}
