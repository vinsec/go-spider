package request

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// isDownload depend if the page need to be downloaded
type Request struct {
	*http.Request
	retryTimes int
	isDownload bool
	depth      int
}

func NewRequest(method string, urlStr string, isDownload bool, depth int, data url.Values) (*Request, error) {
	if urlStr == "" {
		return nil, errors.New("url cannot be empty")
	}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	return &Request{
		Request:    req,
		retryTimes: 0,
		isDownload: isDownload,
		depth:      depth,
	}, nil
}

func (r *Request) IsDownload() bool {
	return r.isDownload
}

func (r *Request) Depth() int {
	return r.depth
}

func (r *Request) RetryTimes() int {
	return r.retryTimes
}

func (r *Request) Url() string {
	return r.Request.URL.String()
}

func (r *Request) GetRequest() *http.Request {
	return r.Request
}

func (r *Request) SetIsDownload(isdownload bool) {
	r.isDownload = isdownload
}

func (r *Request) Valid() bool {
	if r.Url() == "" || r.Request.Method == "" {
		return false
	}
	return true
}

func (r *Request) String() string {
	return fmt.Sprintf(
		"Request Method: %s, Url: %s, Retry Times: %d, Depth: %d, Download: %t",
		r.Request.Method,
		r.Request.URL.String(),
		r.retryTimes,
		r.depth,
		r.isDownload)
}
