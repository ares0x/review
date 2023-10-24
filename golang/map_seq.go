package main

import "fmt"

var m = map[string]int{
	"ares":10,
	"echo":8,
	"kobe":4,
}

func main() {
	rangeMap()
	rangeMapSeq() //有序
}

func rangeMap() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func rangeMapSeq() {
	arr := make([]string, 0, len(m))
	for k := range m {
		arr = append(arr, k)
	}
	for _, v := range arr {
		fmt.Printf("key:%s,value:%d",v, m[v])
	}
}