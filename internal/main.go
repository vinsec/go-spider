// Copyright 2022 vinsec. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

/*
modification history
--------------------
2022/07/17 21:30:20, by vinsec, create
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
)

import (
	"github.com/vinsec/go-spider/config"
	"github.com/vinsec/go-spider/downloader"
	"github.com/vinsec/go-spider/parser"
	"github.com/vinsec/go-spider/spider"
	"github.com/vinsec/go-spider/util"
)

var (
	displayVersion bool
	confFile       string
	logDir         string
)

func init() {
	flag.BoolVar(&displayVersion, "v", false, "display spider version then exit")
	flag.StringVar(&confFile, "c", "../conf/spider.conf", "set config file path")
	flag.StringVar(&logDir, "l", "../log/", "set log directory")
}

// main the function where execution of the program begins
func main() {
	flag.Parse()
	flag.Usage = util.DisplayHelpMenu
	if displayVersion {
		spider.DisplaySpiderVersionExit()
	}
	confFile = path.Clean(confFile)
	cfg, err := config.LoadConfigFromFile(confFile)
	if err != nil {
		fmt.Println("Load Config File Err: " + err.Error())
		os.Exit(1)
	}

	logDir = path.Clean(logDir)
	if !util.IsDirExist(logDir) {
		if ok, err := util.Mkdir(logDir); !ok {
			fmt.Println("Create Log Dir Err: ", err.Error())
			os.Exit(1)
		}
	}
	util.InitialLogger("sSpider", "INFO", logDir, true, "midnight", 3)

	runtime.GOMAXPROCS(runtime.NumCPU())
	mainSpider := spider.NewSpider(cfg.Spider.ThreadCount, cfg.Spider.MaxDepth, cfg.Spider.CrawlInterval)

	var origins map[string]bool
	origins, err = util.GetSeedFromFile(cfg.Spider.UrlListFile, false)
	if err != nil {
		util.Logger.Error(err)
		os.Exit(1)
	}
	if len(origins) <= 0 {
		util.Logger.Warn("urls empty: ", cfg.Spider.UrlListFile)
		os.Exit(1)
	}
	mainSpider.AppendOrigins(origins)

	p := parser.NewParser(cfg.Spider.TargetUrl)
	mainSpider.AddParser(p)

	d := downloader.NewDownloader(cfg.Spider.CrawlTimeout, cfg.Spider.OutputDirectory)
	mainSpider.AddDownloader(d)

	mainSpider.Start()
}
