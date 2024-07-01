package util

import (
	"errors"
	"fmt"
)

const helpMenu = `
Usage:	mSpider:	./mSpider [options]

Options:
 -c="../conf/spider.conf" : set spider config file
 -l="../log/"             : set log directory
 -v                       : display spider version then exit

Example:
 ./mSpider -c ../conf/spider.conf -l ../log/
`

var ErrHelpMenu = errors.New("help menu displayed")

func DisplayHelpMenu() error {
	fmt.Print(helpMenu + "\n")
	return ErrHelpMenu
}
