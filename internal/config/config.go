package config

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
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
		return false, errors.New("CrawlInterval must greater than zero")
	}
	if c.Spider.CrawlTimeout <= 0 {
		return false, errors.New("CrawlTimeout must greater than zero")
	}
	if c.Spider.TargetUrl == "" {
		return false, errors.New("TargetUrl empty")
	}
	if c.Spider.ThreadCount <= 0 {
		return false, errors.New("ThreadCount must greater than zero")
	}
	return true, nil
}

// LoadConfigFromFile load config file from local
func LoadConfigFromFile(filePath string) (*Config, error) {
	var conf Config
	err := gcfg.ReadFileInto(&conf, filePath)
	if err != nil {
		return nil, err
	}

	// Override config values with environment variables if they exist
	conf.Spider.UrlListFile = getEnv("SPIDER_URL_LIST_FILE", conf.Spider.UrlListFile)
	conf.Spider.OutputDirectory = getEnv("SPIDER_OUTPUT_DIRECTORY", conf.Spider.OutputDirectory)
	conf.Spider.MaxDepth = getEnvAsInt("SPIDER_MAX_DEPTH", conf.Spider.MaxDepth)
	conf.Spider.CrawlInterval = getEnvAsInt("SPIDER_CRAWL_INTERVAL", conf.Spider.CrawlInterval)
	conf.Spider.CrawlTimeout = getEnvAsInt("SPIDER_CRAWL_TIMEOUT", conf.Spider.CrawlTimeout)
	conf.Spider.TargetUrl = getEnv("SPIDER_TARGET_URL", conf.Spider.TargetUrl)
	conf.Spider.ThreadCount = getEnvAsInt("SPIDER_THREAD_COUNT", conf.Spider.ThreadCount)

	configDir := filepath.Dir(filePath)
	conf.Spider.UrlListFile = resolvePath(configDir, conf.Spider.UrlListFile)
	conf.Spider.OutputDirectory = resolvePath(configDir, conf.Spider.OutputDirectory)

	if _, err := conf.Check(); err != nil {
		return nil, err
	}

	return &conf, nil
}

func resolvePath(basePath, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}
	return filepath.Join(basePath, targetPath)
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the function returns its value.
// If the variable is not present, it returns the default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt is a helper function that retrieves the value of the environment variable as an integer.
func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
