pano-on 例子参考

- 生成镜像

```
    #!/usr/bin/env bash #!/usr/bin/env bash #在不同的系统上提供了一些灵活性
    set -ex   # 打开调试回响模式，并且在出错时推出。
    # CGO_ENABLED=0 GOOS=linux Linux下编译Mac平台的64位可执行程序
    # go build -installsuffix cgo 在软件包安装的目录中增加后缀标识，以保持输出与默认版本分开
    # -o参数，只被允许在编译一个单独包时使用，并且强制将编译的可执行文件或object写入户名的文件名中。
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o panoon ./cmd/server/main.go 

    # docker build -f Dockerfile_PATH .
    也可以加入-t参数指定构建后的镜像名称、标签
    （如：docker build -t nginx:old -f Dockerfile_PATH .）
    docker build -t vcloud/new-panoon:latest -f ./Dockerfile .
    docker login harbor.develop.pano-on.com -u cpj -p Tjb@123456
    docker tag vcloud/new-panoon:latest harbor.develop.pano-on.com/tjb/new-panoon:latest
    docker push harbor.develop.pano-on.com/tjb/new-panoon:latest
    docker rmi $(docker images -f "dangling=true" -q)
    docker images | grep vcloud/new-panoon
```

- 拉取镜像
```
From
Usage: FROM [image name]
DockerFile第一条必须为From指令,指定引用的镜像。如果同一个DockerFile创建多个镜像时，可使用多个From指令（每个镜像一次）

拉取上传exp:
 拉取运行镜像
 #!/usr/bin/env bash

set -x
docker rm -f new_panoon
docker rmi $(docker images --filter "dangling=true" -q --no-trunc) 2>/dev/null
docker login harbor.develop.pano-on.com -u cpj -p Tjb@123456
docker pull harbor.develop.pano-on.com/tjb/new-panoon:latest
docker run -d --name new_panoon -p 9003:9003 \
-v /var/log/vcloud/new_panoon:/usr/share/panoon/logs:rw \
-v /etc/vcloud/new_panoon:/usr/share/panoon/conf:ro \
harbor.develop.pano-on.com/tjb/new-panoon:latest
docker ps | grep new_panoon

```
===========
上传镜像

```
#!/usr/bin/env bash
set -ex

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o panoon ./cmd/server/main.go 

docker build -t vcloud/new-panoon:latest -f ./Dockerfile .
docker login harbor.develop.pano-on.com -u cpj -p Tjb@123456
docker tag vcloud/new-panoon:latest harbor.develop.pano-on.com/tjb/new-panoon:latest
docker push harbor.develop.pano-on.com/tjb/new-panoon:latest
docker rmi $(docker images -f "dangling=true" -q)
docker images | grep vcloud/new-panoon
```


