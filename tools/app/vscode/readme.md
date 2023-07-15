- 1. vscode 下载 [🔗](https://code.visualstudio.com/Download)

- 2. 插件安装(建议使用 VPN 安装)

- 3. 设置-配置调整（settings.json）

  - 3.1 设置代理
    ```
       settings.json 增加配置：
       "http.proxy": "http://127.0.0.1:7890",
       "http.proxyStrictSSL": false
       终端输入：
       git config  http.sslVerify false
       git config  http.proxy http://127.0.0.1:7890
    ```

- 4. launch.json 微调整, 用于代码运行断点调试
     ```
     {
             "name": "Launch Package",
             "type": "go",
             "request": "launch",
             "mode": "auto",
             "program": "${fileDirname}",
             "cwd": "${workspaceFolder}",
             "env": {},
             "args": []
         }
     ```
- 5. Go 版本建议 1.19+ [🔗](https://golang.google.cn/dl/)

- 6. Remote 远程开发
  - 6.1 Remote Development 相关插件安装
  - 6.2 点击左下角 SSH -> Remote SSH -> + Add New SSH HOST.. -> 页面 启动 打开文件夹
  - 6.3 本地部分插件无法使用需要在远程服务器上进行安装
 
- 7. setting.json 配置 [🔗](./setting.json)
- 插件集合
  - VsCode+WSL+Docker 开发环境构建指南
    -  配置
        ``` https://blog.csdn.net/hjb2722404/article/details/120738062 ```
    -  连接容器
       ``` https://blog.csdn.net/weixin_40641725/article/details/105512106 ```
    -  基础容器地址
       ``` https://hub.docker.com/r/dingms/ucas-bdms-hw-u64-2019/tags ```
       
  - vscode-pets 在 vscode 里养宠物  使用方式: 1. Ctrl+Shift+P 2. Start pet coding session
  - Chinese (Simplified) (简体中文) Language Pack for Visual Studio Code
  - Draw.io Integration draw格式文件 编辑
  - Log File Highlighter 日志高亮
  - open in browser 在浏览器中打开页面
  - styled-components-snippets 样式调整
  - vscode-icons 文件类型图标
  - Clang-Format proto 文件格式化 1. 编写clang-format配置文件 2. settings 中配置proto3, file.autoSave
  - code-translator 单词翻译
  - Git History 文件的历史情况和修改情况 & GitLens — Git supercharged
  - Makefile Tools Makefile工具扩展
  - markdownlint
  - Path Intellisense 路径自动补全
  - Prettier - Code formatter 代码格式化
  
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - x
  - 

     

[应用集合](../readme.md)
