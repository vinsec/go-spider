package downloader

import (
	"github.com/vinsec/go-spider/request"
	"github.com/vinsec/go-spider/util"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {

	url := "https://www.github.com"
	req := request.NewRequest("GET", url, false, util.SEED_START_DEPTH, nil)

	var timeout uint
	timeout = 2
	outputDir := ""
	d := NewDownloader(int(timeout), outputDir)
	p, err := d.Download(req)
	if err != nil {
		t.Error("http download failed!")
	}

	strBody := p.StringRespBody()
	if !strings.Contains(strBody, "github") {
		t.Error("download html page failed!")
	}
}
