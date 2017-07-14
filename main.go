package main

import "sync"

var (
	// Normal maps Go 1.8
	nMapIntInt map[int]int
	nMapStrInt map[string]int
	nMapIntStr map[int]string
	nMapStrStr map[string]string

	// Concurrency maps Go 1.9
	cMap sync.Map
)

func init() {
	nMapIntInt = make(map[int]int)
	nMapStrInt = make(map[string]int)
	nMapIntStr = make(map[int]string)
	nMapStrStr = make(map[string]string)

	cMap = sync.Map{}
}

func NMapIntInt_Read(k int) int {
	return nMapIntInt[k]
}

func NMapIntInt_Write(k, v int) {
	nMapIntInt[100] = v
}

func NMapStrInt_Read(k string) int {
	return nMapStrInt[k]
}

func NMapStrInt_Write(k string, v int) {
	nMapStrInt["test"] = v
}

func NMapIntStr_Read(k int) string {
	return nMapIntStr[k]
}

func NMapIntStr_Write(k int, v string) {
	nMapIntStr[100] = v
}

func NMapStrStr_Read(k string) string {
	return nMapStrStr[k]
}

func NMapStrStr_Write(k, v string) {
	nMapStrStr["test"] = v
}

func CMap_Read(k int) interface{} {
	v, _ := cMap.Load(k)
	return v
}

func CMap_Write(k, v int) {
	cMap.Store(100, v)
}
