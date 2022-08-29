package parser

import (
	"regexp"
	"strings"
)
import (
	"golang.org/x/net/html"
)

import (
	"github.com/vinsec/go-spider/request"
	"github.com/vinsec/go-spider/response"
	"github.com/vinsec/go-spider/util"
)

type Parser struct {
	target *regexp.Regexp
}

func NewParser(exp string) *Parser {
	target := regexp.MustCompile(exp)
	return &Parser{target: target}
}

func (p *Parser) Parse(resp *response.Response) error {
	if resp == nil || !resp.IsParse() {
		return nil
	}

	//get response body with string
	bodyString := resp.StringRespBody()
	if bodyString == "" {
		return nil
	}
	util.Logger.Info(">>>>>> Parsing <<<<<< ", resp)

	//read bodyString to a new reader
	bodyReader := strings.NewReader(bodyString)

	//trans the charset of the reader to UTF-8
	strUTFBody, err := util.TransCharsetUTF8(resp.RespHeader("Content-Type"), bodyReader)
	if err == nil {
		bodyReader = strings.NewReader(strUTFBody)
	}

	//
	root, err := html.Parse(bodyReader)
	if err != nil {
		return err
	}
	p.ParseByNode(resp, root)
	return nil
}

func (p *Parser) ParseByNode(resp *response.Response, n *html.Node) {
	if n.Type == html.ElementNode {
		//range all the node's Attr, to find all the links
		for _, arr := range n.Attr {
			if arr.Key == "href" || arr.Key == "src" {
				refUrl := arr.Val
				if strings.HasPrefix(arr.Val, "javascript") {
					refUrl = util.GetUrlFromJs(refUrl)
				}
				absUrl, err := util.TransUrlFromRelToAbs(resp.URL(), refUrl)
				if err != nil {
					util.Logger.Warn("relative url absolute failed: ", refUrl, err)
					continue
				}

				req := request.NewRequest("GET", absUrl, false, resp.Depth()+1, nil)
				req.Header.Add("Referer", resp.Url())
				//in this spider, we only download file with suffix *.html|*.htm
				if p.target.MatchString(absUrl) {
					req.SetIsDownload(true)
				}
				//saving sub page to the resp's commingCrawlReq list
				resp.AddComingCrawlReq(req)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.ParseByNode(resp, c)
	}
}
