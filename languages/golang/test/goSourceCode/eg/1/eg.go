package main

import (
	"log"
	_ "net/http/pprof"
	"runtime"
)

var intMap map[int]struct{}
var cnt = 8192

func main() {

	printMemStats()

	initMap()
	runtime.GC()
	printMemStats()

	// log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		// time.Sleep(time.Millisecond)
		delete(intMap, i)
		// runtime.GC()
		// printMemStats()
	}
	// log.Println(len(intMap))

	runtime.GC()
	printMemStats()
	// fmt.Println("v: ", v)
	intMap = nil
	printMemStats()
	runtime.GC()
	// // fmt.Println("v: ", v)
	printMemStats()
}

func initMap() {
	intMap = make(map[int]struct{}, cnt)

	for i := 0; i < cnt; i++ {
		intMap[i] = struct{}{}
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
	/*
		Alloc：     当前堆上对象占用的内存大小;
		TotalAlloc：堆上总共分配出的内存大小;
		Sys：       程序从操作系统总共申请的内存大小;
		NumGC：     垃圾回收运行的次数。
	*/
}
