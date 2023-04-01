- 1. vscode 下载 [🔗](https://code.visualstudio.com/Download)
    

- 2. 插件安装(建议使用VPN安装)

  - Chinese (Simplified) (简体中文) L （适用于 VS Code 的中文（简体）语言包）
  - Go （Go for Visual Studio Code）
  - json （美化 json 插件）
  - open in browser （右击选中默认的浏览器打开 html 文件）
  - Path Intellisense （路径提示插件）
  - Prettier - Code formatter （自动格式化插件）
  - Styled-Components Snippets
  - vscode-icons （可以控制 vscode 中的文件管理的树目录显示图标）
  - vscode-styled-components （语法提示高亮）
  - Git Extension Pack
  - GitLens — Git supercharged
  - Comment Translate
  - ctr+shift+p go tools install && install/update tools
  

- 3. 设置-配置调整（settings.json）
  -  3.1 设置代理
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
  - 6.2 点击左下角 SSH   -> Remote SSH -> + Add New SSH HOST.. -> 页面 启动 打开文件夹
  - 6.3 本地部分插件无法使用需要在远程服务器上进行安装
