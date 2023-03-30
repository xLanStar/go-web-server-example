package test

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

var (
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 20)
		},
	}
)

func TestSlice(t *testing.T) {

	varr := bufferPool.Get()
	arr := varr.([]byte)
	fmt.Println(arr)
	f, _ := os.OpenFile("test.txt", os.O_RDONLY, 0644)
	len, _ := f.Read(arr)
	f.Close()
	fmt.Println(arr[:len])
	fmt.Println(string(arr[:len]))
	bufferPool.Put(varr)

	varr = bufferPool.Get()
	arr = varr.([]byte)
	fmt.Println(arr)
	f, _ = os.OpenFile("test2.txt", os.O_RDONLY, 0644)
	len, _ = f.Read(arr)
	f.Close()
	fmt.Println(arr[:len])
	fmt.Println(string(arr[:len]))
	bufferPool.Put(varr)
}
