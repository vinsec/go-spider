package config

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {

	//just check path.Clean()
	rawPath := "../../conf/spider.conf_test"
	cleanPath := path.Clean(rawPath)
	if cleanPath != "../../conf/spider.conf_test" {
		t.Error("clean function not ok")
	}

	//just check filepath.Abs() filepath.Dir()
	currentPath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	parentPath := filepath.Dir(filepath.Dir(currentPath)) + "/conf"
	absPath, err := filepath.Abs(filepath.Dir(cleanPath))

	if err != nil || absPath != parentPath {
		t.Error(absPath + err.Error())
	}

	conf, err := LoadConfigFromFile(cleanPath)
	if err != nil {
		t.Error(err)
	}

	if conf.Spider.UrlListFile != "../../data/seed" {
		t.Error("test LoadConfigFromFile failed")
	}

	if conf.Spider.OutputDirectory != "../../output/" {
		t.Error("testLoadConfigFromFile failed")
	}



}
