- ```
  高效的存储介质　内存条(１５０ ns)
  优良的底层数据结构　hashTable(O(1))
  高效的网络IO模型 (epoll, kqueue, select)
  高效的线程模型 (多线程&单线程)
  ```

- hashtable

  ```
  hash(key)%hashtable.size 存储值指针
  
  hash冲突 entry 中 存有 next 指针，头插法，链表形式进行存储
  
  redis扩容 hashtable 成倍形式增长，旧数据空间 渐进式 搬到新数据空间上 rehash
  ```

- string

  ```
  sdshdr5、 sdshdr8、 sdshdr16、sdshdr32、sdshdr64、 (优化)
  alloc 值调整增长优化 free -> alloc 表示剩余多少长度可用
  sdshdr5 => flags(1 btye) buf(...btye) => flags(Type(3bit), Len(5bit))
  
  ```

  

-  set

  ```
  SADD key v1 v2 v3 v4 ...
  object encoding key 
  底层两种数据类型(intset, hastable)，纯int型能够进行自动去重排序(制定长度内)，hastable 是无序数据结构
  intset 转为 hastable 的两种情况
  1. 元素个数大于set-max-intset-entries（值可配）
  2. 元素无法整形表示
  
  intset 底层就是数组 （int16, int32, int64 根据位数不同类型，节省空间）判断是否存在使用的二分查找发
  
  ```

  

- hash

  ```
  hash 数据结构底层实现为一个字典（dict）,也是RedisDB 用来存储kv的数据结构，当数据量比较小，或者单个元素表较小时，底层用ziplist存储，数据大小和元素数量阈值可以通过参数进行配置
  
  hash-max-ziplist-entries 512
  ziplist 元素个数超过512时，将改为hashtable编码
  
  hash-max-ziplist-value 64
  单个元素大小超过64Byte时，将改为hashtable编码
  ```

- zset

  ```
   ZSet 是有序的，自动去重的集合数据类型，ZSet 数据结构底层实现为字典（dict）+ 跳表（skiplist）,当数据比较少时，用ziplist编码结构存储
   
   zset-max-ziplist-entries 128
   元素个数超过128时，将用skiplist编码
   
   zset-max-ziplist-value 64
   元素个数超过64Byte，将用skiplist
   
   dict 快速索引数据 计算得分 存储索引
   skiplist 存储有序数据
   
   reamLevel 概率 p^level-1 * (1-p)  | p的n-1次方*（1-p）| 造成情况 高层节点少，低层节点多
   eg:
   level = 3;
   	第一次： random < zksiplist_p 概率为p
   	第二次： random < zksiplist_p 概率为p
   	第三次： random > zksiplist_p 概率为（1-p）
   	
  ```

  

- 