package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "fuck.h"
*/
import "C"

import (
	"os/exec"
	"strings"
)

const defaultCLIPath = "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"

var urlListener chan string = make(chan string)

func main() {
	go C.RunApp()
	url := <-urlListener

	flag := "--folder-uri"
	if strings.HasSuffix(url, ".code-workspace") {
		flag = "--file-uri"
	}
	cmd := exec.Command(defaultCLIPath, flag, url)
	cmd.Run()
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
