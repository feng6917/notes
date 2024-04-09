#### 目录

[设置时区](#设置时区)

[apt常用命令](#apt常用命令)

[文件&文件夹常用命令](#文件&文件夹常用命令)

[性能相关](#性能相关)
[](#)
[](#)

#### 设置时区

- 临时设置  

  ```
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

  ```

- 同步设置

  ```
    root@ubuntu1804:~# apt install ntpdate

    root@ubuntu1804:~# date 
    Tue Dec 13 15:12:07 CST 2022
    
    root@ubuntu1804:~# date -d '-1 day'
    Mon Dec 12 15:12:22 CST 2022
    
    root@ubuntu1804:~# ntpdate ntp.aliyun.com
    13 Dec 15:12:36 ntpdate[15040]: adjust time server 203.107.6.88 offset 0.002588 sec
    
    root@ubuntu1804:~# date 
    Tue Dec 13 15:12:40 CST 2022  #同步后的时间
  ```

[目录](#目录)

#### apt常用命令

- apt

  ```
    列出所有可更新的软件清单命令：sudo apt update

    升级软件包：sudo apt upgrade

    列出可更新的软件包及版本信息：apt list --upgradeable

    升级软件包，升级前先删除需要更新软件包：sudo apt full-upgrade

    安装指定的软件命令：sudo apt install <package_name>

    安装多个软件包：sudo apt install <package_1> <package_2> <package_3>

    更新指定的软件命令：sudo apt update <package_name>

    显示软件包具体信息,例如：版本号，安装大小，依赖关系等等：sudo apt show <package_name>

    删除软件包命令：sudo apt remove <package_name>

    清理不再使用的依赖和库文件: sudo apt autoremove

    移除软件包及配置文件: sudo apt purge <package_name>

    查找软件包命令： sudo apt search <keyword>

    列出所有已安装的包：apt list --installed

    列出所有已安装的包的版本信息：apt list --all-versions
  ```

[目录](#目录)

#### 文件&文件夹常用命令

- 处理目录的常用命令

  ```
    ls（英文全拼：list files）: 列出目录及文件名
    cd（英文全拼：change directory）：切换目录
    pwd（英文全拼：print work directory）：显示目前的目录
    mkdir（英文全拼：make directory）：创建一个新的目录
    rmdir（英文全拼：remove directory）：删除一个空的目录
    cp（英文全拼：copy file）: 复制文件或目录
    rm（英文全拼：remove）: 删除文件或目录
    mv（英文全拼：move file）: 移动文件与目录，或修改文件与目录的名称
  ```

- Linux 文件内容查看

  ```
    # 查询服务器中大于1G的文件
    find / -type f -size +1G

    # 查询服务器中大于1G的文件及属性信息
    find / -type f -size +1G  -print0 | xargs -0 ls -l

    # 查询大文件并排序
    find / -type f -size +1G  -print0 | xargs -0 du -h | sort -nr

    cat  由第一行开始显示文件内容
    tac  从最后一行开始显示，可以看出 tac 是 cat 的倒着写！
    nl   显示的时候，顺道输出行号！
    more 一页一页的显示文件内容
    less 与 more 类似，但是比 more 更好的是，他可以往前翻页！
    head 只看头几行
    tail 只看尾巴几行

    # tail 命令可用于查看文件的内容，有一个常用的参数 -f 常用于查阅正在改变的日志文件。
    tail -f --since 1s logs.txt
  ```

- 拷贝文件

  ```
    # 拷贝本地文件到服务器 可拷贝多个
    scp test root@VM2:/backup

    # 拷贝服务器文件到本地 -r 递归文件夹
    scp -r root@192.168.163.130:/root/ /root

    # 拷贝服务器到另一台服务器
    scp root@192.168.163.128:/root/test3 root@192.168.163.130:/backup/
  ```

- 压缩解压缩文件

  ```
    # 压缩文件(带不带gz取决于带不带z)
    tar -zcvf test.tar.gz ./xxx
    # 解压缩
    tar -zxvf ./test.tar.gz

  ```  

- grep

  ```
    grep -v 排除内容
    grep -B 显示匹配行和之前num行
    grep -A 显示匹配行和之后num行
    grep -C 显示匹配行和前后num行
    grep --color=auto 匹配字符串加色显示
    grep -n 打印行号
    grep -i 不区分大小写
    grep -w 匹配单词
    grep -E 即egrep使用扩展正则表达式
    grep -e 匹配多个模式
    grep -c 匹配到的行数
  ```

[目录](#目录)

#### 性能相关

- 查看进程&端口

  ```
    # 查找指定进程格式
    ps -ef | grep 进程关键字

    # 杀死进程 强杀 -9
    kill -9 PID

    # 查询端口
      1. 查询所有
      netstat -ntlp
      2. 查询指定 
      losf -i:xxxx
      netstat -tunlp |grep xxxx
  ```

- 日志输出

  ```
  # shell中可能经常能看到：echo log > /dev/null 2>&1 命令的结果可以通过%>的形式来定义输出
    /dev/null ：代表空设备文件
    >  ：代表重定向到哪里，例如：echo "123" > /home/123.txt
    1  ：表示stdout标准输出，系统默认值是1，所以">/dev/null"等同于"1>/dev/null"
    2  ：表示stderr标准错误
    &  ：表示等同于的意思，2>&1，表示2的输出重定向等同于1
    2>&1 & 将这个（标准输出&错误输出）任务放到后台去执行
  ```

- 空间占用

  ```  
  # 查询文件 空间占用
  [root@k8s-master-147 glusterVol]# du -sh model/
  997M model/
  # 空间使用率
  [root@k8s-master-7 ~]# df ./myz/ -h
  Filesystem               Size  Used Avail Use% Mounted on
  /dev/mapper/centos-root   50G   41G  9.7G  81% /
  ```

[目录](#目录)
  
[命令集合](./readme.md)
