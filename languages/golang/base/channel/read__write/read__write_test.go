package read__write

import "testing"

func TestCh1CloseRead(t *testing.T) {
	/*
		关闭chan 一般放在写入数据的地方；
		chan 一直可以读取数据，读取到的数据跟chan 关闭前是否存在元素有关，存在正常读取，第二个返回值 true, 反之 false;
		无缓存chan 同步, 有缓存chan 异步。
	*/
}
