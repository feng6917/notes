#### test

- 基本命令

  ```
  测试 该路径下所有测试
  go test -v ./ -count=1

  测试 指定某个测试
  go test -v -run TestSum ./

  -v 输出完整的测试结果
  -count= 1 清除缓存
  -run 它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。Testxxx 指定测试某个函数

  ```

- 全局测试 前后操作 TestMain
- 单个测试 前后操作 setupTestSum
- 单个 case 测试 TestSum
- 多个 case 测试 TestSum2
- 基准测试 BenchmarkSum
- 并行基准测试 BenchmarkSumParallel
- 示例函数 ExampleSum

- 参考链接：
  - [Go 语言基础之单元测试](https://www.liwenzhou.com/posts/Go/unit-test/)

[Golang 目录](../../readme.md)
