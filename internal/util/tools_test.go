package util

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	dummyFile := "./testFile"
	if IsFileExist(dummyFile) {
		t.Error("test IsFileExist() failed")
	}
	realFile := "./tools.go"
	if !IsFileExist(realFile) {
		t.Error("test IsFileExist() failed")
	}
}

func TestIsDirExist(t *testing.T) {
	pwd, _ := os.Getwd()
	realDir := filepath.Dir(pwd)
	if !IsDirExist(realDir) {
		t.Error("test IsDirExist() failed")
	}
	dummyPath := realDir + "/dummy"
	if IsDirExist(dummyPath) {
		t.Error("test IsDirExist() failed")
	}
}

func TestTransUrlFromRelToAbs(t *testing.T) {
	baseUrl, _ := url.Parse("http://www.baidu.com")
	refUrl := "/testurl"
	resUrl, err := TransUrlFromRelToAbs(baseUrl, refUrl)
	if err != nil || resUrl != "http://www.baidu.com/testurl" {
		t.Error("test TransUrlFromRelToAbs() failed " + err.Error())
	}
}
