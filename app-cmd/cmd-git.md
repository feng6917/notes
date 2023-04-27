
#### Git 常用命令

```
# 把要提交的文件的信息添加到暂存区中
git add

# 列出本地的所有分支，当前所在分支以 "*" 标出
git branch

# 创建新分支，新的分支基于上一次提交建立
git branch <分支名>

# 删除指定的本地分支
git branch -d <分支名称>

# 强制删除指定的本地分支
git branch -D <分支名称>

# 切换到已存在的指定分支
git checkout <分支名称>

# 创建并切换到指定的分支，保留所有的提交记录
git checkout -b <分支名称>

# 默认在当前目录下创建和版本库名相同的文件夹并下载版本到该文件夹下
git clone <远程仓库的网址>

# 指定本地仓库的目录
git clone <远程仓库的网址> <本地目录>

# -b 指定要克隆的分支，默认是master分支
git clone <远程仓库的网址> -b <分支名称> <本地目录>

# 把暂存区中的文件提交到本地仓库中并添加描述信息
git commit -m "<提交的描述信息>"

# 修改上次提交的描述信息
git commit --amend

# 查看配置信息 --local：仓库级，--global：全局级，--system：系统级
git config <--local | --global | --system> -l

# 查看当前生效的配置信息
git config -l

# 配置提交记录中的用户信息 --global 全局
git config --global user.name <用户名>
git config --global user.email <邮箱地址>

# 比较当前文件和暂存区中文件的差异，显示没有暂存起来的更改
git diff

# 比较两个分支之间的差异
git diff <分支名称> <分支名称>

# 将远程仓库所有分支的最新版本全部取回到本地
git fetch <远程仓库的别名>

# 初始化本地仓库，在当前目录下生成 .git 文件夹
git init

# 打印所有的提交记录
git log

# 打印指定数量的最新提交的记录
git log -<指定的数量>

# 把指定的分支合并到当前所在的分支下，并自动进行新的提交
git merge <分支名称>

# 重命名文件或者文件夹
git mv <源文件/文件夹> <目标文件/文件夹>

# 从远程仓库获取最新版本。
git pull

# 把本地仓库的分支推送到远程仓库的指定分支
git push <远程仓库的别名> <本地分支名>:<远程分支名>

# 删除指定的远程仓库的分支
git push <远程仓库的别名> :<远程分支名>
git push <远程仓库的别名> --delete <远程分支名>

# 将 HEAD 的指向改变，撤销到指定的提交记录，文件未修改
git reset <commit ID>

# 生成一个新的提交来撤销某次提交，此次提交之前的所有提交都会被保留
git revert <commit ID>

# 移除跟踪指定的文件/文件夹，并从本地仓库的文件夹中删除
git rm -r <文件路径>

# 查看本地仓库的状态
git status

# 打印所有的标签
git tag

# 添加轻量标签，指向提交对象的引用，可以指定之前的提交记录
git tag <标签名称> [<commit ID>]

# 切换到指定的标签
git checkout <标签名称>

# 删除指定的标签
git tag -d <标签名称>

# 将指定的标签提交到远程仓库
git push <远程仓库的别名> <标签名称>

# 将本地所有的标签全部提交到远程仓库
git push <远程仓库的别名> –tags

# 本地修改数据抛弃
git stash
git stash drop

# 设置代理
git config --global http.proxy socks5://127.0.0.1:1080
git config --global https.proxy socks5://127.0.0.1:1080


# 创建 submodule
git submodule add <submodule_url>

# 获取 submodule
一种方式是在克隆主项目的时候带上参数 --recurse-submodules

一种可行的方式是，在当前主项目中执行：
git submodule init & git submodule update

```

#### 本地合并commitID
```
# 查看前10个commit
git log -10
# 从版本库恢复文件到暂存区，不改动工作区的内容
git reset --soft 295ac3b842b4ecb6eff1c9954a281a4606a8bc84	# 别人改的commitID
# add已经跟踪的文件
git add -u
# 提交
git commit -m "修改信息"
# 强制push以替换远程仓的commitID
git push --force
```

#### 多分支操作
```
# 添加新的工作树
git worktree add ../project-name-branch-name (-b) branch-name
# 查看工作树列表
git worktree list
# 删除工作树
git worktree remove project-name-branch-name
```


#### Git 建议使用提交规范

```
<type>(<scope>): <subject>
build(package.json): 修改typescript版本到3.4.1

选择改动类型 (<type>)
填写改动范围 (<scope>)
写一个精简的描述 (<subject>)
```



- type 

  ```
  type为必填项，用于指定commit的类型，约定了feat、fix两个主要type，以及docs、style、build、refactor、revert五个特殊type，其余type暂不使用。
  ```

  - 主要type

    > feat: 增加新功能
    >
    > fix: 修复bug

  - 特殊type

    > docs 只改动了文档相关的内容
    >
    > style 不影响代码含义的改动，例如去掉空格、改变缩进、增删分号
    >
    > build 构造工具的或者外部依赖的改动，例如webpack、npm
    >
    > refactor 代码重构时使用
    >
    > revert 执行 git revert打印的message

  - 暂不使用type

    > test 添加测试或者修改现有测试
    >
    > perf 提高性能的改动
    >
    > ci 与CI(持续集成服务)有关的改动
    >
    > chore 不修改src或者test的其余修改

 - scope

   ```
   scope也为必填项，用于描述改动的范围，格式为项目名/模块名，例如：node-pc/common rrd-h5/activity，而we-sdk不需指定模块名。如果一次commit修改多个模块，建议拆分成多次commit，以便更好追踪和维护。
   ```

   
