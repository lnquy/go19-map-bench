package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// Normal map
	nMapIntInt map[int]int
	nMapStrInt map[string]int
	nMapIntStr map[int]string
	nMapStrStr map[string]string

	// Concurrency map Go 1.9
	sMap sync.Map
)

func init() {
	nMapIntInt = make(map[int]int)
	nMapStrInt = make(map[string]int)
	nMapIntStr = make(map[int]string)
	nMapStrStr = make(map[string]string)
	sMap = sync.Map{}

	// Maps to test main func
	nMap = make(map[int]int)
	sMap2 = sync.Map{}
}

// Benchmark functions
func NMapIntInt_Read(k int) int {
	return nMapIntInt[k]
}

func NMapIntInt_Write(k, v int) {
	nMapIntInt[k] = v
}

func NMapStrInt_Read(k string) int {
	return nMapStrInt[k]
}

func NMapStrInt_Write(k string, v int) {
	nMapStrInt[k] = v
}

func NMapIntStr_Read(k int) string {
	return nMapIntStr[k]
}

func NMapIntStr_Write(k int, v string) {
	nMapIntStr[k] = v
}

func NMapStrStr_Read(k string) string {
	return nMapStrStr[k]
}

func NMapStrStr_Write(k, v string) {
	nMapStrStr[k] = v
}

func SMap_Read(k int) interface{} {
	v, _ := sMap.Load(k)
	return v
}

func SMap_Write(k, v int) {
	sMap.Store(k, v)
}

// --------------------------------------------------
// Test concurrency write time for 1 million elements
var (
	nMap map[int]int
	mux  sync.RWMutex
)

var (
	sMap2 sync.Map
)

var (
	n = 1000000 // Number of elements to insert into map
	g = 10      // Number of goroutines
)

func fillNormalMap(w <-chan int, r chan<- bool) {
	for i := range w {
		mux.Lock()
		nMap[i] = i
		mux.Unlock()
	}
	r <- true
}

func testNormalMapConcurrency() {
	defer execTime("Normal map", time.Now())
	w := make(chan int, n)
	res := make(chan bool, g)
	for i := 0; i < g; i++ {
		go fillNormalMap(w, res)
	}

	for i := 0; i < n; i++ {
		w <- i
	}
	close(w)

	for i := 0; i < g; i++ {
		<-res
	}
}

func fillSyncMap(w <-chan int, r chan<- bool) {
	for i := range w {
		sMap2.Store(i, i)
	}
	r <- true
}

func testSyncMapConcurrency() {
	defer execTime("Sync map 1.9", time.Now())
	w := make(chan int, n)
	res := make(chan bool, g)
	for i := 0; i < g; i++ {
		go fillSyncMap(w, res)
	}

	for i := 0; i < n; i++ {
		w <- i
	}
	close(w)

	for i := 0; i < g; i++ {
		<-res
	}
}

func execTime(n string, t time.Time) {
	fmt.Printf("%s executed in %v\n", n, time.Now().Sub(t))
}

func main() {
	testNormalMapConcurrency()
	fmt.Printf("Normal map size: %v\n", len(nMap))

	testSyncMapConcurrency()
	var sMapLen int
	sMap2.Range(func(_, _ interface{}) bool {
		sMapLen++
		return true
	})
	fmt.Printf("Sync map size: %v\n", sMapLen)
}
