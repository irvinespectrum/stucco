// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

var m map[int64]int64
var wg sync.WaitGroup
var lock = sync.RWMutex{}

func TestHello(t *testing.T) {

	m = make(map[int64]int64)

	runtime.GOMAXPROCS(2)
	wg.Add(3)
	go A()
	go A()
	go A()

	for i := 0; i < 25; i++ {
		id := GetID()
		fmt.Println(id)
	}

	wg.Wait()
	fmt.Println(len(m))
}

func A() {
	for i := 0; i < 25; i++ {
		id := GetID()
		fmt.Println(id)
	}
	wg.Done()
}

func GetID() int64 {
	rs, err := http.Get("http://localhost:8080/?id=test")
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	i64, err := strconv.ParseInt(bodyString, 10, 64)
	writeMap(i64)
	return i64
}

func writeMap(i int64) {
	lock.Lock()
	defer lock.Unlock()
	if m[i] != 0{
		panic("duplicate id")
	}
	m[i] = i
}
