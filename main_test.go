package main

import (
	"fmt"
	"testing"
)

var (
	tmpInt int
	tmpStr string
	tmpItf interface{}
)

func BenchmarkNMapIntInt_Write(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		NMapIntInt_Write(tmp, tmp)
	}
}

func BenchmarkNMapIntInt_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpInt = NMapIntInt_Read(i % 1000)
	}
}

func BenchmarkNMapStrInt_Write(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		NMapStrInt_Write(fmt.Sprintf("%s", tmp), i)
	}
}

func BenchmarkNMapStrInt_Read(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		tmpInt = NMapStrInt_Read(fmt.Sprintf("%s", tmp))
	}
}

func BenchmarkNMapIntStr_Write(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		NMapIntStr_Write(tmp, fmt.Sprintf("%s", tmp))
	}
}

func BenchmarkNMapIntStr_Read(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		tmpStr = NMapIntStr_Read(tmp)
	}
}

func BenchmarkNMapStrStr_Write(b *testing.B) {
	var tmp string
	for i := 0; i < b.N; i++ {
		tmp = fmt.Sprintf("%s", i%1000)
		NMapStrStr_Write(tmp, tmp)
	}
}

func BenchmarkNMapStrStr_Read(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		tmpStr = NMapStrStr_Read(fmt.Sprintf("%s", tmp))
	}
}

func BenchmarkCMap_Write(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		SMap_Write(tmp, tmp)
	}
}

func BenchmarkCMap_Read(b *testing.B) {
	var tmp int
	for i := 0; i < b.N; i++ {
		tmp = i % 1000
		tmpItf = SMap_Read(tmp)
	}
}
