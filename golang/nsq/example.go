package main

func main() {
	// 消费订阅及心跳参考官方，后续补充吧
	// https://github.com/nsqio/go-nsq
	// docker 部署nsq操作 参考链接： https://www.cnblogs.com/jssyjam/p/11546233.html
	// 操作如下（ip替换为外网地址）：
	// - nsqlookupd
	// docker run --name lookupd -p 4160:4160 -p 4161:4161 -d nsqio/nsq /nsqlookupd
	// - nsqd
	// docker run --name nsqd -p 4150:4150 -p 4151:4151 -d nsqio/nsq /nsqd --broadcast-address=ip --lookupd-tcp-address=ip:4160
	// - nsqadmin
	// docker run -d --name nsqadmin -p 4171:4171 nsqio/nsq /nsqadmin --lookupd-http-address=ip:4161

}
