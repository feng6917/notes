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
 
- 7. setting.json 配置
     ```
       {
          "[python]": {
            "editor.formatOnType": true
          },
          "settingsSync.ignoredExtensions": ["alefragnani.project-manager"],
          "go.lintFlags": [
            // 提示等级
            "-min_confidence=.8"
          ],
          // 检查工具，默认staticcheck
          "go.lintTool": "staticcheck",
          // 输入提示上下文
        
          // https://github.com/golang/vscode-go/blob/master/docs/settings.md#uidiagnosticanalyses
        
          "go.editorContextMenuCommands": {
            "toggleTestFile": true,
        
            "addTags": true,
        
            "removeTags": false,
        
            "fillStruct": true,
        
            "testAtCursor": true,
        
            "testFile": false,
        
            "testPackage": false,
        
            "generateTestForFunction": true,
        
            "generateTestForFile": false,
        
            "generateTestForPackage": false,
        
            "addImport": true,
        
            "testCoverage": true,
        
            "playground": true,
        
            "debugTestAtCursor": true,
        
            "benchmarkAtCursor": false
          },
          // git设置，vscode中设置中文后，鼠标悬浮会有中文提示
        
          "git.untrackedChanges": "separate",
        
          "git.alwaysShowStagedChangesResourceGroup": true,
        
          // 工作区颜色设置，可自定义主题，此处的配置优先级为最高，切换主题不会改变这里的配置
        
          "workbench.colorCustomizations": {
            // 鼠标选择的文字的颜色
            "editor.selectionHighlightBackground": "#fd2605",
        
            // ctrl + f 时的搜索时选中的文字颜色
            "editor.findMatchBackground": "#fa2404"
          },
        
          // -count=1 清除缓存
          "go.testFlags": [
            "-ldflags",
            // "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore",
            "-v",
            "-count=1"
          ],
        
          // 自动保存
          "files.autoSave": "onFocusChange",
          "editor.formatOnSave": true,
          "go.toolsManagement.autoUpdate": true,
          // 设置代理
          // "http.proxy": "http://127.0.0.1:7890",
          // "http.proxyStrictSSL": false
          "[proto3]": {
            "editor.defaultFormatter": "xaver.clang-format"
          },
          "window.zoomLevel": 1,
          "go.testTimeout": "120s"
        }

     ```
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
