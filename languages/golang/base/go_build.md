#### go build 构建
---
- GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
- GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm
---
- Mac Mac下编译Linux, Windows平台的64位可执行程序：
    - Linux
      ``` 
      CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appName main.go
      ```
    - Windows
      ```
      CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o appName main.go
      ```
- Linux Linux下编译Mac, Windows平台的64位可执行程序：
    - Mac
      ```
      CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o appName main.go
      ```
    - Windows
      ```
      CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o appName main.go
      ```
- Windows
    - Mac
      ```
        SET CGO_ENABLED=0
        SET GOOS=darwin
        SET GOARCH=amd64
        go build -o appName main.go
      ```
    - Linux
      ```
        SET CGO_ENABLED=0
        SET GOOS=linux
        SET GOARCH=amd64
        go build -o appName main.go
      ```
