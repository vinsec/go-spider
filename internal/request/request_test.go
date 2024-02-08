package request

import (
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	data := url.Values{}
	data1 := url.Values{}
	newReq := NewRequest("GET", "www.baidu.com", true, 1, data)
	if newReq.Url() != "www.baidu.com" {
		t.Error("test Url() failed")
	}
	if !newReq.IsDownload(){
		t.Error("test IsDownload() failed")
	}

	newReqWithNilUrl := NewRequest("GET", "xxx", true, 1, data1)
	if !newReqWithNilUrl.Valid(){
		t.Error("test Valid() failed")
	}
}
