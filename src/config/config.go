package config

import (
	"errors"
	"path/filepath"
)

import (
	"github.com/go-gcfg/gcfg"
)

import (
	"github.com/vinsec/go-spider/util"
)

type Config struct {
	Spider struct {
		UrlListFile     string // 种子文件路径
		OutputDirectory string // 下载目录
		MaxDepth        int    // 最大抓取深度
		CrawlInterval   int    // 抓取间隔
		CrawlTimeout    int    // 抓取超时
		TargetUrl       string // 目标文件正则
		ThreadCount     int    // 抓取routine数
	}
}

func (c *Config) Check() (bool, error) {
	if !util.IsFileExist(c.Spider.UrlListFile) {
		return false, errors.New("UrlListFile not exist: " + c.Spider.UrlListFile)
	}
	if !util.IsDirExist(c.Spider.OutputDirectory) {
		return false, errors.New("OutputDirectory not exist: " + c.Spider.OutputDirectory)
	}
	if c.Spider.CrawlInterval < 0 {
		return false, errors.New("OutputDirectory must greater than zero")
	}
	if c.Spider.CrawlTimeout <= 0 {
		return false, errors.New("OutputDirectory must greater than zero")
	}
	if c.Spider.TargetUrl == "" {
		return false, errors.New("TargetUrl empty")
	}
	if c.Spider.ThreadCount <= 0 {
		return false, errors.New("ThreadCount must greater than zero")
	}
	return true, nil
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	var conf Config
	err := gcfg.ReadFileInto(&conf, filePath)
	if err != nil {
		return nil, err
	}

	var configDir string
	configDir, err = filepath.Abs(filepath.Dir(filePath))
	if !util.IsFileExist(conf.Spider.UrlListFile) {
		urlListFile := configDir + "/" + conf.Spider.UrlListFile
		if util.IsFileExist(urlListFile) {
			conf.Spider.UrlListFile = urlListFile
		}
	}
	if !util.IsDirExist(conf.Spider.OutputDirectory) {
		outputDir := configDir + "/" + conf.Spider.UrlListFile
		if util.IsDirExist(outputDir) {
			conf.Spider.OutputDirectory = outputDir
		}
	}
	_, err = conf.Check()
	if err != nil {
		return nil, err
	}
	return &conf, nil

}
