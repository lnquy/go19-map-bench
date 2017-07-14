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

func BenchmarkNMapIntInt_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpInt = NMapIntInt_Read(i)
	}
}

func BenchmarkNMapIntInt_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NMapIntInt_Write(i, i)
	}
}

func BenchmarkNMapStrInt_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpInt = NMapStrInt_Read(fmt.Sprintf("%s", i))
	}
}

func BenchmarkNMapStrInt_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NMapStrInt_Write(fmt.Sprintf("%s", i), i)
	}
}

func BenchmarkNMapIntStr_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpStr = NMapIntStr_Read(i)
	}
}

func BenchmarkNMapIntStr_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NMapIntStr_Write(i, fmt.Sprintf("%s", i))
	}
}

func BenchmarkNMapStrStr_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpStr = NMapStrStr_Read(fmt.Sprintf("%s", i))
	}
}

func BenchmarkNMapStrStr_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NMapStrStr_Write(fmt.Sprintf("%s", i), fmt.Sprintf("%s", i))
	}
}

func BenchmarkCMap_Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpItf = CMap_Read(i)
	}
}

func BenchmarkCMap_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CMap_Write(i, i)
	}
}
