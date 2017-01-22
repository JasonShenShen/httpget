package main

import (
	"fmt"
	"net/http"
	"strconv"
	//"reflect"
	"runtime"
	"strings"
)

var fin = make(chan string)
var limits = make(chan int, 1000)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func doit(url string) {
	limits <- 1
	defer func() {
		<-limits
	}()
	httpget(url)
}

func httpget(url string) {
	resp, err := http.Get(url)
	defer func() {
		fmt.Printf("url is %s, goroutine id is %d\n", url, GoID())
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		fmt.Println(url, err)
	} else {
		fmt.Println(url, resp.StatusCode)
	}
	fin <- url
}

func main() {

	var httplist [1000]string

	for i := 0; i < len(httplist); i++ {
		httplist[i] = "http://www.baidu.com/s?wd=search" + strconv.Itoa(i+1)
		//httplist[i] = "http://192.168.6.150:2003/"
		go doit(httplist[i])
	}

	//等待结束
	for i := 0; i < len(httplist); i++ {
		<-fin
	}
}
