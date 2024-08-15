#### go-micro

- 引入外部框架 Gin 生成 Web Api

  ```
  ginRouter := gin.Default()
  server := web.NewService(
   ...
   web.Handler(ginRouter)
   ...
  )
  ```

- 服务注册 快速注册到 Consul 中

  ```
  consulReg := consul.NewRegistry(
   registry.Addrs("xxxxx:xxx")
  )
  server := web.NewService(
   ...
   web.Registry(consulReg)
   ...
  )
  ```

- 服务发现 selector (随机、轮询)

  ```
  consulReg := consul.NewRegistry(
   registry.Addrs("xxxxx:xxx")
  )
  getService, err := consulReg.GetSevice("xxxService")
  deal(err)
  
  next := selector.Random(getService)
  node, err := next()
  deal(err)
  ```

- 内置参数启动 注册多服务

  ```
  go run xxx.go --server_address :8080
  ```

- 服务调用 插件方式调用 http、rpc (proto 格式引入 tag json protoc-go-inject-tag 修改tagName)

- gin ctx 服务传参 中间件封装

  ```
  func InitMildreware(s service) gin.HandlerFunc() {
   return func(ctx *gin.Context) gin.HandlerFunc() {
    ctx.Keys = make(map[string]interface{})
    ctx.Keys["service"] = s
    ctx.Next()
   }
  }
  ```

- warpper (装饰器 - logger) , broker (熔断器 - hystrix 超时，服务降级 断言重写response )

- micro 工具箱

  ```
  # 查看服务
  micro --registry consul --registry_address xxx list services 
  # 调用服务
  micro --restistry consul --registry_address xxx call service_name func_name request_body
  
  # 为 rpc 服务创建 网关
  ```

- micro 弃用consul 使用ETCD ，1. k8s 使用etcd 主流 2. 使用consul 一些问题
