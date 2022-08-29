package parser

import (
	"github.com/vinsec/go-spider/request"
	"github.com/vinsec/go-spider/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestParser(t *testing.T) {
	data := make(url.Values)
	u := "http://www.baidu.com"
	req := request.NewRequest("GET", u, true, 0, data)
	client := &http.Client{}

	page := response.NewResp(req)

	//do a real request
	resp, err := client.Do(req.Request)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	//apply all the attrs to page
	page.SetRespBody(content)
	page.SetRespStatus(resp.StatusCode)
	page.SetRespHeader(resp.Header)
	page.SetIsParse(true)

	parser := NewParser(".*.(htm|html)$")
	err = parser.Parse(page)
	if err != nil {
		t.Error(err)
	}
	var urlLists []string
	for _, r := range page.GetComingCrawlReq() {
		if r.Url() == "http://map.baidu.com" || r.Url() == "http://news.baidu.com" {
			urlLists = append(urlLists, r.Url())
		}
	}
	if len(urlLists) != 2 {
		t.Error("test Parse() error")
	}

}
