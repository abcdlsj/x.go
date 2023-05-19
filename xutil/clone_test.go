package xutil

import "testing"

type Data struct {
	Name string
	Age  int
}

func BenchmarkCloneStruct(b *testing.B) {
	data := Data{Name: "Alice", Age: 26}
	var newData Data
	for i := 0; i < b.N; i++ {
		Clone(&newData, &data)
	}
}

func BenchmarkCloneStr(b *testing.B) {
	data := "Alice"
	var newData string
	for i := 0; i < b.N; i++ {
		Clone(&newData, &data)
	}
}

func BenchmarkCloneInt(b *testing.B) {
	data := 49
	var newData int
	for i := 0; i < b.N; i++ {
		Clone(&newData, &data)
	}
}
