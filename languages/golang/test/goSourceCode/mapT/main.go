package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

/*
   map 是一个无序的键值对的集合。
   Go 语言中的 map 是一个引用类型，这意味着它是一个值，可以作为函数参数传递，
   可以作为结构体字段的值，甚至可以作为切片或者其他 map 的值。

   map 的定义格式为：make(map[KeyType]ValueType)

   map 删除内存不会立即释放, 下次GC时也不会释放  设置为Nil后 会在下次GC时进行释放
*/

/*
   在 Go 语言中，`map` 是一种内置的数据类型，用于存储键值对。`map` 的底层实现是基于哈希表（hash table）的，这意味着它提供了快速的查找、插入和删除操作。

	以下是 `map` 的底层实现的一些关键点：

	1. **哈希表**：`map` 的底层实现是基于哈希表的。哈希表是一种数据结构，它使用哈希函数将键转换为哈希值，然后根据哈希值将键值对存储在适当的桶（bucket）中。

	2. **桶**：每个桶是一个链表，用于存储具有相同哈希值的键值对。当哈希冲突发生时，新的键值对将被添加到链表的末尾。

	3. **负载因子**：`map` 的负载因子是当前存储的键值对数量与桶数量的比值。当负载因子超过一定阈值时，`map` 会进行扩容，即创建一个新的、更大的哈希表，并将所有键值对重新哈希到新的哈希表中。

	4. **并发访问**：`map` 不是并发安全的，如果你在多个 goroutine 中同时访问和修改 `map`，可能会导致数据竞争和不一致的问题。如果你需要在多个 goroutine 中安全地访问和修改 `map`，可以使用 `sync.Map` 或者在访问和修改 `map` 时使用互斥锁（`sync.Mutex`）。

	需要注意的是，`map` 的底层实现是复杂的，涉及到哈希函数的选择、哈希冲突的处理、扩容的策略等。如果你需要深入了解 `map` 的底层实现，可以参考 Go 语言的源代码。
*/

func main() {
	// fmt.Println("Hello, World!")
	deleteT()
} // https://golang.org/doc/go1.4#maps

func deleteT() {
	m := make(map[int]int)
	m[1] = 1000000
	m[2] = 2000000000
	m[3] = 3000000
	m[4] = 40000000000
	fmt.Printf("before val count:%d, addr : %v, size: %v \r\n", len(m), &m, unsafe.Sizeof(m))
	delete(m, 1)
	time.Sleep(time.Second)
	runtime.GC()
	printMemStats("删除一条数据")
	fmt.Printf("after val count:%d, addr : %v, size: %v \r\n", len(m), &m, unsafe.Sizeof(m))
	fmt.Printf("after timeSleep val count:%d, addr : %v, size: %v \r\n", len(m), &m, unsafe.Sizeof(m))
	delete(m, 2)
	delete(m, 3)
	delete(m, 4)
	time.Sleep(time.Second)
	runtime.GC()
	printMemStats("删除所有数据")
	fmt.Printf("after delete all val count:%d, addr : %v, size: %v \r\n", len(m), &m, unsafe.Sizeof(m))
	m = nil
	printMemStats("删除所有数据后置为nil")
	fmt.Printf("after set nil val count:%d, addr : %v, size: %v \r\n", len(m), &m, unsafe.Sizeof(m))
	time.Sleep(time.Second)
	runtime.GC()
	printMemStats("GC after set nil")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
