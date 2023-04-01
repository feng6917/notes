package main

import (
	"channel/readWithWrite"
	"channel/simple"
	"channel/sortExec"
	"channel/sortExec/single"
)

func main() {
	mReadWithWrite()
}

func mSimple() {
	// 简单示例
	// 没有缓冲区，同步读写情况下，没有接收者导致死锁
	simple.SimpleA()
	// 设置缓冲区，数据写入缓冲区，chan 关闭后可以正常读取， 但是关闭后不能写入
	simple.SimpleA1()
	// 并发读取，一边写，一边读
	simple.SimpleA2()
}

func mSortExec() {
	/*
		场景：并发模拟吃饭场景，分为三个步骤，1，洗手，2，做饭，3，吃饭，依次进行。
	*/
	//A1()

	// 单个chan实现, 通过判断内部值
	single.Single()
	sortExec.SortExecA()
}

/*
关闭chan 一般放在写入数据的地方；
chan 一直可以读取数据，读取到的数据跟chan 关闭前是否存在元素有关，存在正常读取，第二个返回值 true, 反之 false;
无缓存chan 同步, 有缓存chan 异步。
*/
func mReadWithWrite() {

	//Ch1WriteAndRead()

	//Ch2WriteAndRead()

	//Ch1CloseRead()

	readWithWrite.Ch2CloseRead()
}
