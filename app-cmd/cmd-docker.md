```
    # 运行bash应用
    docker run -it xxxx /bin/bash

    # 查看镜像信息
    docker iamges / docker image ls 

    # 添加标签 (标签起到引用或快捷方式的作用)
    docker tag destTagImage sourceImage

    # 获取镜像详细信息 -f (返回形式：json)
    docker inspect imageID

    # 搜索远端仓库中共享的镜像
    docker search TERM （--automated=false 仅显示自动创建的镜像/ --n0-trunc=false 输出信息不截断显示/ --stars=0 指定仅显示评价为指定星级以上的镜像）
    ep: docker search mysql

    # 删除镜像(-f 强行删除)
    docker rmi imageID

    # 查看本机上存在的所有容器
    docker ps -a

    # 创建镜像三种形式
    基于已有镜像的容器创建/基于本地模板导入/基于Dockerfile创建
            
    存入和载入镜像

    docker save -o dst.tar(指定输出压缩文件) source1-image:tag source2-image:tag ...

    docker load --input source-image.tar /docker load < source-image.tar


    # 上传镜像
    docker push NAME[:TAG]

    # 创建容器
    docker create 

    # 启动容器
    docker start

    # docker run 运行时后台所进行的标注操作：
    . 检查本地是否攒在指定的镜像，不存在就从公有仓库下载
    . 利用镜像创建并启动一个容器
    . 分配一个文件系统，并在只读的镜像层外挂载一层可读写层
    . 从宿主主机配置一个IP地址给容器
    . 执行用户指定的应用程序
    . 执行完毕后容器被终止

    # 启动一个bash终端，允许用户进行交互
    docker run -it xxx-image /bin/bash
    -t 让docker分配一个伪终端并绑定到容器的标准输入上
    -i 让容器的标注输入保持打开
    退出容器 Ctrl+C / exit

    # 让docker容器在后台以守护态（Daemonized）形式运行 -d
    docker run -d xxx-image /bin/bash

    # 终止运行中的容器 docker stop
    docker ps -a -q 查看处于终止状态的容器的id信息

    # 将一个运行态的容器终止，重新启动
    docker restart xxx-image

    # 进入容器
    docker exec -ti xxx-image /bin/bash

    # 删除处于终止状态的容器
    docker rm [options] container [container...]
    -f --force=false 强行终止并删除一个运行中的容器
    -l --link=false 删除容器的连接，但保留容器
    -v --volumes=false 删除容器挂在的数据卷

    # 导出容器
    docker export container > xxx.tar

    # 导入容器
    cat xxx.tar | docker import xxx.tar

    # 映射到指定地址的指定端口
    -P 是允许外部访问容器需要暴露的端口
    docker -p ip:port:containerPort

    # 查看映射端口配置
    docker port 查看当前映射的端口设置

    # 容器有自己的内部网络和IP地址（使用docker inspect+ 容器id可以获取所有的变量值）

    # dockerfile 由一行行命令语句组成，支持以#开头的注释行
    第一行必须指定基于的基础镜像
    FROM ...

    # 维护者信息
    MAINTAINER docker_user docker_user@email.com

    # 容器启动时执行命令
    CMD /usr/sbin/nginx

    # FROM 格式为FROM image /image:tag
    第一条指令必须为FROM指令。并且，如果在同一个Dockerfile中创建多个镜像时，可以使用多个FROM指令（每个镜像一次）

```