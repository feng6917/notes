工作中多次遇到导出EXCEL文件的接口，能否可以写成统一格式进行调用呢，最简单的方式可以利用结构体的属性，即反射来实现，搞一下。

在 golang 中最简单的方式来获取一个结构体的所有tag，如下；
```
import reflect
type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(name, s.Tag)
	}
}
reflect.TypeOf方法获取对象的类型，之后NumField()获取结构体成员的数量。 通过Field(i)获取第i个成员的名字。 再通过其Tag 方法获得标签。
```

关于生成Excel的思考？
```
    首先我们清楚最基本的Excel包含表头及表数据两个基本元素，表头为每一列元素的阐述说明，表示这一列数据是什么？而每一列数据则是阐述说明下的具体数据。
    生成Excel最基本的操作，写入表头，写入数据。在写入数据时，我们可能会对要写入的数据进行处理，例如数据是枚举，例如数据要进行特殊处理，这些都需要根据实际特殊处理。
    能够通过结构体的tag进行处理呢？答案是肯定，在解析结构体tag时，根据自定义表头tag字段名称获取表头名称，根据自定义枚举tag字段名称处理数据值，根据自定义方法tag字段名称处理方法数据值。
   
```
具体如何生成Excel的代码如下：[仅供参考](https://github.com/feng6917/blog/blob/main/golang/excelTag
)