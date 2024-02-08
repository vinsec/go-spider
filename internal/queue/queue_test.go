package queue

import (
	"github.com/vinsec/go-spider/request"
	"net/url"
	"testing"
)

func TestQueue(t *testing.T) {
	data := make(url.Values)
	queue := NewTaskQueue()
	req := request.NewRequest("GET", "www.baidu.com", false, 1, data)
	queue.Push(req)
	if queue.Count() != 1 {
		t.Error("test Count() failed")
	}
	//test hashtable
	req_ := request.NewRequest("POST", "www.baidu.com", true, 1, data)
	queue.Push(req_)
	if queue.Count() != 1 {
		t.Error("test queue uniq failed")
	}
	popReq, err := queue.Pop()
	if err != nil || popReq.Url() != "www.baidu.com" {
		t.Error("test Pop() failed")
	}
}
