#### doc

*缺什么*

```
  etcd
  redis
  架构理解
  分布式服务及分布式锁
  红黑树
  消息队列
  服务注册与发现
  熔断器 限流器
  线上问题排查
  tidb
  linux 性能内存排查
```

```
  gin
  go-micro
  broker 熔断器
```

- golang 开发新手常犯的五十个错误
  - 1. 不允许出现未使用的变量 & import ?
  - 2. 变量声明的方式有几种, 分别是什么 ? 是否能够使用短变量重复声明 ?
  - 3. 幽灵变量是什么？特点？
  - 4. 字符串是否可以修改，是否可以为nil?
  - 5. 数组,切片, map, slice, chan 函数传参时如何传递, 数组和切片有什么区别？
  - 6. 从不存在key的map中取值返回？map返回的数据时一致的么？
  - 7. interface nil 是否可以比较？
  - 8. for 中大聪明（go func()、append()、defer file close()）
  - 9. defer 后面跟的是什么？有什么特点？
  - 10. 变量内存分析，堆，栈是什么，逃逸分析是什么？
    - 逃逸分析 <https://geektutu.com/post/hpg-escape-analysis.html>
  - 11. 零切片, nil切片和空切片是否一样？
  - 12. 字符串转成byte数组，会发生内存拷贝么？
    - <https://blog.csdn.net/lsoxvxe/article/details/132427676>
  - 13. 拷贝大切片与小切片代价是否一致？
  - 14. map 是否可以不进行初始化，初始化长度区别？承载多大，iterator 是否安全，能不能一边del一边range?线程安全的map怎么实现
    - <https://juejin.cn/post/7215587423685378085>
  - 15. 翻转含有中文、数字、英文字母的字符串
  - 16. map触发扩容的时机，满足什么条件时扩容？扩容策略是什么?
  - 17. new() 和 make() 的区别
    - [answer](./17.md)
  - 18. go struct能不能比较？
    - [answer](./18.md)
  - 19. go map如何顺序读取？
    - <https://segmentfault.com/a/1190000043834586>
  - 20. go中怎么实现set
  - 21. Golang 有没有 this 指针？指针是什么？
    - <https://blog.csdn.net/fly910905/article/details/105989267>
  - 22. Golang 中的引用类型包含哪些?
    - <https://blog.csdn.net/luduoyuan/article/details/135396996>
  - 23. string, byte, rune ? 利用unsafe.Pointer和reflect包可以实现很多禁忌的黑魔法，但这些操作对GC并不友好。最好不要尝试。
    - [黑魔法](./23.md)
  - 24. for select时，如果通道已经关闭会怎么样，如果有一个case, 一个defualt? 如果只有一个case呢？
    - <https://mp.weixin.qq.com/s/TuuLYgvIkwREDLkALqTMXA>
  - 25. select可以用于什么？
    - <https://learnku.com/articles/82805>
  - 26. context 的用途？
    - <https://segmentfault.com/a/1190000040917752>
  - 27. 怎样避免内存逃逸？
    - <https://developer.baidu.com/article/detail.html?id=3316082>
  - 28. goroutine 泄露的概念及常见原因？(至少说出四种)
    - <https://blog.csdn.net/Ws_Te47/article/details/135521647>
  - 29. 内存泄漏的场景及解决方案？（至少说出四种）
    - <https://blog.csdn.net/Ws_Te47/article/details/135521647>
    - [io pprof图](./leak/io/pprof.png)
  - 30. sync.Pool的适用场景
    - <https://developer.baidu.com/article/detail.html?id=3229046>
  - 31. sync.Map 优缺点，使用场景
    - <https://blog.csdn.net/lsoxvxe/article/details/132427824>
  - 35. go 主协程等待子协程执行完毕再执行方法几种方式
    - <https://blog.csdn.net/dqz_nihao/article/details/124904807>
  - 36. go chan 有缓冲和无缓冲的区别
    - <https://www.runoob.com/note/43083>
  - 37. 携程间的通信方式
    - <https://blog.csdn.net/qq_17199495/article/details/125787317> toTest
  - 38. go 中chan的底层原理
    - <https://blog.csdn.net/qq_58244272/article/details/136895661> eg

  - 39. 读写锁底层是怎么实现的？
    - <https://www.cnblogs.com/peteremperor/p/14097633.html>
    - video <https://www.bilibili.com/video/BV1rg411B71e/?spm_id_from=pageDriver&vd_source=7d32ad5a1a541e44326e50415ffd9907>
  - 40. golang的CSP思想?
    - <https://www.jianshu.com/p/36e246c6153d>  
  - 41. 能说说uintptr和unsafe.Pointer的区别么？
  - 42. reflect(反射包)如何获取字段tag? 为什么json包不能导出私有变量的tag?
  - 43. 进程，线程，携程？
    <https://blog.csdn.net/EDDYCJY/article/details/116141654>
  - 44. 垃圾回收的过程是怎么样的？什么是写屏障、混合写屏障，如何实现？(参考前记录)
  - 45. GMP 模型 协程之间是怎么调度的
    <https://zboya.github.io/post/go_scheduler/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io>
  - 46. gc的stw是怎么回事?
  - 47. 利用golang特性，设计一个QPS为500的服务器(针对问题需要发散思考，不要盲目的解题)
    <https://blog.csdn.net/micl200110041/article/details/82013032>  
  - 48. 必须要手动对齐内存的情况
    <https://juejin.cn/post/7082332804922966023>  
  - 49. 堆内存和栈内存分配方式， Go协程的栈内存管理（后续再说）
    go 语言设计与实现
    <https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-stack-management/>
  - 50. go 常用的优化手段？[<-](./50.md)
  - 51. 怎样访问私有变量？(没有意义)
    <https://www.jianshu.com/p/7b3638b47845>
  - 52. 类型断言的大概实现过程？
    <https://segmentfault.com/a/1190000039894161>
  - 53. 接口原理及使用场景?
    <https://juejin.cn/post/7171288417324498980?searchId=20240626104405564347837F7AF32BAC9F>
  - 54. 为什么小对象多了会造成 gc 压力?
  - 55. 闭包怎么实现的,闭包的主要应用场景？
  - 56. 两次 GC 周期重叠会引发什么问题， Gc的时机有哪些，能手动触发么？
    <https://blog.csdn.net/kevin_tech/article/details/122613350>
  - 57. Goroutinue 什么时候会被挂起？
    <https://blog.csdn.net/asd1126163471/article/details/124893098>  
  - 58. Data Race 问题怎么检测？怎么解决?
    <https://blog.csdn.net/raoxiaoya/article/details/118437969>  
    <https://learnku.com/articles/45279>
  - 59. Golang 触发异常的场景有哪些?
  - 60. net/http包中client如何实现长连接？
    <https://segmentfault.com/a/1190000042631284>
  - 61. rpc 了解
    <https://www.cnblogs.com/sumuncle/p/11554904.html>
  - 62. pb buffer 原理
    <https://blog.csdn.net/dyj5841619/article/details/94717419>

  - 63. socket 是什么？
      <https://golangguide.top/%E8%AE%A1%E7%AE%97%E6%9C%BA%E5%9F%BA%E7%A1%80/%E7%BD%91%E7%BB%9C%E5%9F%BA%E7%A1%80/%E6%A0%B8%E5%BF%83%E7%9F%A5%E8%AF%86%E7%82%B9/socket%E5%88%B0%E5%BA%95%E6%98%AF%E4%BB%80%E4%B9%88%EF%BC%9F.html>

  - 64. tcp 粘包 数据包问题？
  - 65. 既然IP层会分片，为什么TCP层也要分段？
  - 66. 断网了，还能ping通 127.0.0.1 吗？ 为什么？
  - 67. 连接一个IP不存在的主机时，握手过程是怎样的？
  - 68. 代码执行send成功后，数据就发出去了吗？
  - 69. 收到RST， 就一定会断开TCP连接么？
  - 70. 没有accept, 能建立TCP链接么？
  - 71. HTTP 是无状态的吗？需要保持状态的场景应该怎么做？
  - 72. RestFul 是什么？RestFul 请求的 URL 有什么特点？
  - 73. 一次url访问会经历哪些过程
  - 74. TCP 三次握手以及四次挥手的流程。
    <<https://segmentfault.com/a/1190000022082901> 凑合看>
  - 75. TCP的拥塞控制具体是怎么实现的？UDP有拥塞控制吗？
      <https://juejin.cn/post/6981357492466892836>
  - 76. 是否了解中间人劫持原理
      <https://segmentfault.com/a/1190000041047662>
  - 77. TCP 与 UDP 在网络协议中的哪一层，他们之间有什么区别？
  - 78. HTTP 与 HTTPS 有哪些区别？
  - 79. select, poll 和epoll区别
      <https://zhuanlan.zhihu.com/p/629960221>
  - 80. TCP 如何实现数据有序性？
      <<https://www.coonote.com/tcpip-note/tcp-ensures-order-transmission.html> 了解一下 就是排序的具体步骤>
  - 81. TCP长连接和短连接有那么不同的使用场景？
      <<https://www.jianshu.com/p/1cbc522c983d> 了解一下>
  - 82. TIME_WAIT时长，为什么？
  - 83. 什么是零拷贝？
    <https://segmentfault.com/a/1190000044068914>
  - 84. HTTP 简述 HTTP 的 keepalive 的原理和使用场景
    <https://juejin.cn/post/7116843553505935367>
  - 85. Cookie 和 Session 的关系和区别是什么？
  - 89. DNS 查询服务器的基本流程是什么？DNS 劫持是什么？
  －90. 简述对称与非对称加密的概念
  - 91. epoll中的ET和LT模式 ???
  - 92. JWT 的原理和校验机制
  - 93. 介绍下proactor和reactor ???
    <https://cloud.tencent.com/developer/article/1769945>
  - 94. RPC 的调用过程 ???
    <https://cloud.tencent.com/developer/article/1727582>
  - 95. 什么是 ARP 协议？简述其使用场景
    <https://info.support.huawei.com/info-finder/encyclopedia/zh/ARP.html>
  - 96. DDOS 攻击原理，如何防范它？
    <https://segmentfault.com/a/1190000044892118>
  - 97. 如何防止传输内容被篡改？
  - 98. 什么是中间人攻击？如何防止攻击？
    <https://info.support.huawei.com/info-finder/encyclopedia/zh/%E4%B8%AD%E9%97%B4%E4%BA%BA%E6%94%BB%E5%87%BB.html>
  - 99. udp包长度
    <https://www.cnblogs.com/sinferwu/articles/7615276.html>
  - 100. IP为什么要分片
    <https://cloud.tencent.com/developer/article/2335459>
  - 101. 数据包乱序如何处理
  - 102. 什么是 SYN flood，如何防止这类攻击？
    <https://info.support.huawei.com/info-finder/encyclopedia/zh/SYN+Flood.html>
  - 103. websocket 是如何传输的
    <https://cloud.tencent.com/developer/article/1887095>
  - 104. 为什么需要序列化？有什么序列化的方式？
  - 105. 有chunked的时候contentlength是什么样子
    <https://cloud.tencent.com/developer/article/1804395>
  - 106. 如何设计一个可靠的udp
  - 107. TCP 中常见的拥塞控制算法有哪些？
    <https://blog.csdn.net/Hesy_H/article/details/86574003>
  - 108. 如何设置非阻塞
      <https://blog.csdn.net/analogous_love/article/details/134720873>
      <https://www.cnblogs.com/Suzkfly/p/14331287.html>
  - 109. 什么是跨域，什么情况下会发生跨域请求？
      <https://cloud.tencent.com/developer/article/1678220>

  - 111. traceroute 有什么作用？
      <https://www.cnblogs.com/emilyyoucan/articles/7478532.html>
  - 112. TIME_WAIT危害
      <https://www.cnblogs.com/dabai56/p/11201106.html>
  - 114. select什么情况返回0
  - 115. TCP 的 keepalive 了解吗？说一说它和 HTTP 的 keepalive 的区别？
      <https://www.cnblogs.com/shitianming/p/16823530.html>
  - 116. 简述常见的 HTTP 状态码的含义
      <https://www.runoob.com/http/http-status-codes.html>  
  - 200.

  - 暂停  

----

#### 数据库

- 400. 数据库三大范式是什么
- 401. mysql有关权限的表都有哪几个
- 402. MySQL的binlog有有几种录入格式？分别有什么区别？
- 403. mysql有哪些数据类型
      <https://www.cnblogs.com/secsafe/p/12255370.html>
- 404. MySQL存储引擎MyISAM与InnoDB区别
- 405. MyISAM索引与InnoDB索引的区别？
  <https://www.cnblogs.com/balfish/p/8288891.html>
- 406. InnoDB引擎的4大特性
- 407. 什么是索引？优缺点，聚簇索引与非聚簇索引
  <https://developer.aliyun.com/article/841106>
- 408. 索引使用场景
  <https://www.cnblogs.com/ghjbk/p/8384262.html>
- 409. 索引算法有哪些？
- 410. 百万级别或以上的数据如何删除
- 411. 前缀索引
    <https://www.cnblogs.com/itjiangpo/p/14181386.html>
- 412. 什么是最左前缀原则？
    <https://blog.csdn.net/zhizhengguan/article/details/120855011>
- 413. B树和B+树的区别
    <https://www.cnblogs.com/1832921tongjieducn/p/13358264.html>  
- 414. 使用B树的好处, 使用B+树的好处
- 415. Hash索引和B+树所有有什么区别或者说优劣呢?
- 416. 非聚簇索引一定会回表查询吗？
- 417. 什么是数据库事务？
  <https://cloud.tencent.com/developer/article/1816543>
- 418. 对MySQL的锁了解吗，锁的是什么？
  <https://cloud.tencent.com/developer/article/2185599>  
- 419. 按照锁的粒度分数据库锁有哪些？锁机制与InnoDB锁算法  
- 420. MySQL中InnoDB引擎的行锁是怎么实现的？
- 421. 数据库的乐观锁和悲观锁是什么？怎么实现的？
  <https://cloud.tencent.com/developer/article/1920876>
- 422. 为什么要使用视图？什么是视图？
- 423. 视图的优点, 缺点
- 424. 什么是游标？
    <https://www.cnblogs.com/luoshengjie/p/10153933.html>
- 425. 存储过程与函数，存储过程有哪些优缺点？
- 426. 什么是触发器？触发器的使用场景有哪些？MySQL中都有哪些触发器？
- 427. SQL语句主要分为哪几类
      <https://www.cnblogs.com/pingdingshan/p/15934188.html>
- 428. 超键、候选键、主键、外键分别是什么？
    <https://www.cnblogs.com/zhoading/p/10928769.html>  
- 429. SQL 约束有哪几种？
    <https://www.cnblogs.com/willingtolove/p/9215330.html>
- 430. 六种关联查询
- 431. 什么是子查询
    <http://t.csdnimg.cn/78Ip5>
- 432. mysql中 in 和 exists 区别
- 433. mysql中int(10)和char(10)以及varchar(10)的区别
- 434. FLOAT和DOUBLE的区别是什么？
- 435. UNION与UNION ALL的区别？
- 436. 如何定位及优化SQL语句的性能问题？
    <https://cloud.tencent.com/developer/article/1665882>
- 437. SQL的生命周期？
    <https://developer.aliyun.com/article/1391548>
- 438. 大表数据查询，怎么优化
    <https://blog.csdn.net/Y_hanxiong/article/details/136694475>
- 439. 超大分页怎么处理？
    <https://blog.csdn.net/m0_56615376/article/details/137048998>
    <https://blog.csdn.net/QQ727338622/article/details/127224611>
- 440. 慢查询日志怎么看
    <https://www.cnblogs.com/greatsql/p/17011719.html>
- 441. MySQL数据库cpu飙升到500%的话他怎么处理？  
- 442. 分库分表了是怎么做的？分表分库了有什么问题？有用到中间件么？
  <https://www.cnblogs.com/heqiyoujing/p/10924574.html>
- 443. MySQL的复制原理以及流程
  <https://www.cnblogs.com/miracle77hp/p/10208405.html>  
- 444. 读写分离有哪些解决方案？
  <https://cloud.tencent.com/developer/article/1664449>  
- 445. 备份计划，mysqldump以及xtranbackup的实现原理
  <https://www.cnblogs.com/zhoujinyi/p/5789465.html>
- 446. 数据表损坏的修复方式有哪些？
  <https://www.cnblogs.com/ygcphp/p/10282747.html>

  ---

- 100. 进程为什么比线程要快？
- 101. 进程间通信方式有哪些？

  ---

- 200. mysql 与 redis 如何保证双写一致性
  
- 201. 网络及操作系统面试题 <https://golangguide.top/%E8%AE%A1%E7%AE%97%E6%9C%BA%E5%9F%BA%E7%A1%80/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/%E9%9D%A2%E8%AF%95%E9%A2%98.html>

  ---

- 300. 深度解密 Go 语言之 sync.map
    <https://qcrao.com/post/dive-into-go-sync-map/>
    <https://blog.csdn.net/Dong_chongwu/article/details/128862493>

- 301. mutex

- 302. chan 发送及接收数据
