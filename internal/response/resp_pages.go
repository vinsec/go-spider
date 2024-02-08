package response

import (
	"fmt"
	"net/http"
	"net/url"
)

import (
	"github.com/vinsec/go-spider/request"
)

// struct "Response" means downloaed page and its functions
// isParse depend if the page need to be parsed
type Response struct {
	request         *request.Request
	respStatus      int
	respHeader      http.Header
	respCookie      []*http.Cookie
	respBody        []byte
	commingCrawlReq []*request.Request
	isParse         bool
}

func NewResp(req *request.Request) *Response {
	return &Response{request: req}
}

// if 200<= status code <300, marked success
func (r *Response) Success() bool {
	if r.respStatus >= http.StatusOK && r.respStatus < http.StatusMultipleChoices {
		return true
	}
	return true
}

func (r *Response) RespHeader(key string) string {
	return r.respHeader.Get(key)
}

func (r *Response) RespBody() []byte {
	return r.respBody
}

func (r *Response) StringRespBody() string {
	return string(r.respBody)
}

func (r *Response) GetComingCrawlReq() []*request.Request {
	return r.commingCrawlReq
}

func (r *Response) IsParse() bool {
	return r.isParse
}

func (r *Response) URL() *url.URL {
	return r.request.URL
}

func (r *Response) Url() string {
	return r.request.Url()
}

func (r *Response) Depth() int {
	return r.request.Depth()
}

func (r *Response) AddComingCrawlReq(req *request.Request) {
	r.commingCrawlReq = append(r.commingCrawlReq, req)
}

func (r *Response) SetRespBody(body []byte) {
	r.respBody = body
}

func (r *Response) SetRespStatus(status int) {
	r.respStatus = status
}
func (r *Response) SetRespCookies(cookies []*http.Cookie) {
	r.respCookie = cookies
}

func (r *Response) SetRespHeader(header http.Header) {
	r.respHeader = header
}

func (r *Response) SetIsParse(isParse bool) {
	r.isParse = isParse
}

func (r *Response) String() string {
	return fmt.Sprintf("Req: %s, Resp Code: %d, Need Parse: %t",
		r.request.String(),
		r.respStatus,
		r.isParse)
}
