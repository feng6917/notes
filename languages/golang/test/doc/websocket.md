websocket

---

Request
---

accept-encoding: gzip, deflate
accept-language: zh-CN,zh;q=0.9
cache-control: no-cache
connection: Upgrade
host: 10.0.0.96:30003
origin: <http://10.0.0.96:30003>
pragma: no-cache
sec-websocket-extensions: permessage-deflate; client_max_window_bits
sec-websocket-key: M0UWzGSA4O9QhWSLGu46VQ==
sec-websocket-version: 13
upgrade: websocket
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36

---

Response
---

connection: upgrade
date: Thu, 15 Aug 2024 06:23:43 GMT
sec-websocket-accept: 75mV0mbzUpJyay1/2Rfmy3k2+dk=
server: nginx/1.25.3

upgrade: websocke

```
    // 兼容http升级协议 协议互通
    sec-websocket-key -> sec-websocket-accept 
    M0UWzGSA4O9QhWSLGu46VQ== -> 75mV0mbzUpJyay1/2Rfmy3k2+dk=
    base64(Sha1(base64(xxx)+GUID))
```
