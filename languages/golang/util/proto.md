- 参数是否必填

  ```
  1. 插件安装
  2. 拉取校验文件到本地
  	xxx/google/api/*
  	
  3. 构建编写
  	.PHONY: grpc
  	# generate grpc code
  	grpc:
  		protoc --proto_path=. \
  		--proto_path=./third_party \
  		--go_out=paths=source_relative:$(COMPILE_TARGET) \
  		--go-grpc_out=paths=source_relative:$(COMPILE_TARGET) \
  		$(API_PROTO_FILES)
  4. 开发编译
  	(google.api.field_behavior) = REQUIRED # 字段必填
  	make .
  ```

  

- 参数校验

  ```
  1. 插件安装
  go install github.com/envoyproxy/protoc-gen-validate@latest
  
  2. 拉取校验文件到本地
  	xxx/xxx/validate.proto	
  3. 构建编写
  	.PHONY: validate
  	# generate validate code
  	validate:
  		protoc --proto_path=. \
             --proto_path=./third_party \
             --go_out=paths=source_relative:$(COMPILE_TARGET)  \
             --validate_out=paths=source_relative,lang=go:$(COMPILE_TARGET) \
             $(API_PROTO_FILES)
  
  4. 开发编译
  	(validate.rules).uint64.gt = 0 # uint64 大于0
  	make .
  ```

