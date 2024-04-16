- 贴吧自动签到测试
- 操作逻辑：
  - 使用本人登录Cookie，获取到用户贴吧列表，遍历列表进行签到。
- 操作步骤
  - 获取cookie
  - 请求协议获取用户关注贴吧列表
  - 遍历贴吧列表 签到（需要获取tbs）

- 配置文件：init.json

    ```
      {
        "Cookie": "BAIDUID=945E9886C9A948095ED63CF6ACBDCEC0:FG=1; cn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf\""
      }
    ```

- [代码位置](../../languages/golang/test/baidustick/example.go)
