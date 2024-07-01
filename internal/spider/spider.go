package spider

import (
	"fmt"
	"os"
	"strings"
	"time"
)

import (
	"github.com/vinsec/go-spider/downloader"
	"github.com/vinsec/go-spider/manager"
	"github.com/vinsec/go-spider/parser"
	"github.com/vinsec/go-spider/queue"
	"github.com/vinsec/go-spider/request"
	"github.com/vinsec/go-spider/util"
)

const (
	SPIDERVERSION = "v1.0"
)

type Spider struct {
	concurrency   int
	origins       []string
	crawlMaxDepth int
	crawlInterval time.Duration
	spiderManager *manager.RoutineManager
	taskQueue     *queue.TaskQueue
	parser        *parser.Parser
	downloader    *downloader.Downloader
}

func NewSpider(concurrency int, maxDepth int, interval int) *Spider {
	crawInterval := time.Duration(interval) * time.Second
	spiderManager := manager.NewRoutineManager(concurrency)
	taskQueue := queue.NewTaskQueue()
	return &Spider{
		concurrency:   concurrency,
		crawlMaxDepth: maxDepth,
		crawlInterval: crawInterval,
		spiderManager: spiderManager,
		taskQueue:     taskQueue,
	}
}

func (s *Spider) AppendOrigins(urls map[string]bool) {
	for url, isDownload := range urls {
		req, _ := request.NewRequest("GET", url, isDownload, util.SEED_START_DEPTH, nil)
		if req != nil {
			s.addStartReq(req)
		}
	}
}

func (s *Spider) addStartReq(req *request.Request) bool {
	if s.addReq(req) {
		s.origins = append(s.origins, req.Url())
		return true
	}
	return false
}

func (s *Spider) addReq(req *request.Request) bool {
	if req == nil || !req.Valid() {
		return false
	}
	s.taskQueue.Push(req)
	return true
}

func (s *Spider) AddDownloader(downloader *downloader.Downloader) {
	s.downloader = downloader
}

func (s *Spider) AddParser(parser *parser.Parser) {
	s.parser = parser
}

func (s *Spider) checkParserDownloader() bool {
	if s.downloader == nil {
		util.Logger.Error("downloader nil, it needs initial")
		return false
	}
	if s.parser == nil {
		util.Logger.Error("parser nil, it needs initial")
	}
	return true
}

func (s *Spider) Start() {
	ok := s.checkParserDownloader()
	if !ok {
		os.Exit(1)
	}
	util.Logger.Info(s)

	for {
		if s.isSpiderFinish() {
			break
		}

		if s.taskQueue.Empty() {
			continue
		}

		sub := s.spiderManager.GetOne()
		if !sub {
			continue
		}

		go func() {
			defer s.spiderManager.FreeOne()

			for {
				req, err := s.taskQueue.Pop()
				if err != nil {
					util.Logger.Error("get req from queue failed: ", err)
					continue
				}
				if req == nil {
					break
				}
				util.Logger.Info(">>>>>> popping request <<<<<< ", req.Url())

				s.process(req)
				s.sleepInterval()
			}
		}()

	}
}

func (s *Spider) process(req *request.Request) {
	page, err := s.downloader.Download(req)
	if err != nil || page == nil || !page.Success() {
		util.Logger.Error("download req err: ", req, err)
		return
	}

	if page.Depth() >= s.crawlMaxDepth {
		page.SetIsParse(false)
	}
	err = s.parser.Parse(page)
	if err != nil {
		util.Logger.Error("parse resp err: ", page, err)
		return
	}

	for _, req := range page.GetComingCrawlReq() {
		if req.Depth() > s.crawlMaxDepth {
			continue
		}
		s.taskQueue.Push(req)
	}
}

func (s *Spider) isSpiderFinish() bool {
	if s.taskQueue.Count() == 0 && s.spiderManager.Used() == 0 {
		util.Logger.Info("request queue nil, all sub spider are idle, Spider exit.")
		time.Sleep(30 * time.Second)
		return true
	}
	return false
}

func (s *Spider) sleepInterval() {
	time.Sleep(s.crawlInterval)
}

func (s *Spider) String() string {
	return fmt.Sprintf(
		"spider routine num: %d, max depth: %d, crawl interval: %02.2f seconds, start urls: %s.",
		s.concurrency,
		s.crawlMaxDepth,
		s.crawlInterval.Seconds(),
		strings.Join(s.origins, ","))
}

func DisplaySpiderVersionExit() {
	fmt.Printf("spider version %s\n", SPIDERVERSION)
	os.Exit(0)
}
