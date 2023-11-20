- 相同结构体不同类型转换
 
- - 最近开发上遇到一个问题，通过grpcgateway 处理后的int64&uint64类型数据均转换成了字符串类型，本身服务于前端，没有任何问题。但是 项目部署现场后，发现需要两套环境，那么就出现一个问题，经过grpcgateway 处理后的数据类型调整为了字符串类型，与原有类型不匹配，解析就出了问题？

- - 思路：
```
尝试过对数据结果进行包装（笨方法，处理起来麻烦），对数据类型调整过，int64&uint64调整为字符串（笨方法，未尝试 且麻烦，需要过多时间调整 且可能引发更多问题），最直接的方式就是处理字符串转int64|uint64时 进行判断，如果类型不一致，对值进行特殊处理即可。
```
- - 具体操作：
```
1. 自己写方法，通过反射 比对 类型，处理值 进行转换，因时间及对嵌套结构处理不完善问题放弃
2. 在原有方法上进行修改，拉取解析包到本地，一步步排查 解析值报错，调整字符串值 为int类型值即可
```
- - json 包：
```
调整位置：
	json/decode.go/(*decodeState).literalStore
	增添代码：
	if v.Type().Kind() == reflect.Int64 || v.Type().Kind() == reflect.Uint64 {
		s := string(item)
		if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
			item = item[1 : len(item)-1]
		}
	}
```
- - 测试代码：
```
type A struct {
		ID int64
	}

    type B struct {
        ID string
    }

    func main() {
        b := B{ID: "23444"}
        bs, err := json.Marshal(b)
        fmt.Println(string(bs), err)

        var a A
        err = json.Unmarshal(bs, &a)
        if err != nil {
            fmt.Printf("err: %v\r\n", err)
        }
        fmt.Println(a, err)
    }

测试结果：
	{"ID":"23444"} <nil>
	{23444} <nil>
```
