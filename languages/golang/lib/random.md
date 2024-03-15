# Random

random 随机数生成器包，可以生成随机[]bytes, int, string。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/random/random.go](https://github.com/duke-git/lancet/blob/main/random/random.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/random"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [RandBytes](#RandBytes)
-   [RandInt](#RandInt)
-   [RandString](#RandString)
-   [RandUpper](#RandUpper)
-   [RandLower](#RandLower)
-   [RandNumeral](#RandNumeral)
-   [RandNumeralOrLetter](#RandNumeralOrLetter)
-   [RandSymbolChar](#RandSymbolChar)
-   [UUIdV4](#UUIdV4)
-   [RandUniqueIntSlice](#RandUniqueIntSlice)
-   [RandFloat](#RandFloat)
-   [RandFloats](#RandFloats)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="RandBytes">RandBytes</span>

<p>生成随机字节切片</p>

<b>函数签名:</b>

```go
func RandBytes(length int) []byte
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/EkiLESeXf8d)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randBytes := random.RandBytes(4)
    fmt.Println(randBytes)
}
```

### <span id="RandInt">RandInt</span>

<p>生成随机int, 范围[min, max)</p>

<b>函数签名:</b>

```go
func RandInt(min, max int) int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/pXyyAAI5YxD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    rInt := random.RandInt(1, 10)
    fmt.Println(rInt)
}
```

### <span id="RandString">RandString</span>

<p>生成给定长度的随机字符串，只包含字母(a-zA-Z)</p>

<b>函数签名:</b>

```go
func RandString(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/W2xvRUXA7Mi)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //pGWsze
}
```

### <span id="RandUpper">RandUpper</span>

<p>生成给定长度的随机大写字母字符串</p>

<b>函数签名:</b>

```go
func RandUpper(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/29QfOh0DVuh)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //PACWGF
}
```

### <span id="RandLower">RandLower</span>

<p>生成给定长度的随机小写字母字符串</p>

<b>函数签名:</b>

```go
func RandLower(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/XJtZ471cmtI)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandLower(6)
    fmt.Println(randStr) //siqbew
}
```

### <span id="RandNumeral">RandNumeral</span>

<p>生成给定长度的随机数字字符串</p>

<b>函数签名:</b>

```go
func RandNumeral(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/g4JWVpHsJcf)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandNumeral(6)
    fmt.Println(randStr) //035172
}
```

### <span id="RandNumeralOrLetter">RandNumeralOrLetter</span>

<p>生成给定长度的随机字符串（数字+字母)</p>

<b>函数签名:</b>

```go
func RandNumeralOrLetter(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/19CEQvpx2jD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandNumeralOrLetter(6)
    fmt.Println(randStr) //0aW7cQ
}
```

### <span id="RandSymbolChar">RandSymbolChar</span>

<p>生成给定长度的随机符号字符串。</p>

<b>函数签名:</b>

```go
func RandSymbolChar(length int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Im6ZJxAykOm)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandSymbolChar(6)
    fmt.Println(randStr) // 随机特殊字符字符串，例如: @#(_")
}
```

### <span id="UUIdV4">UUIdV4</span>

<p>生成UUID v4字符串</p>

<b>函数签名:</b>

```go
func UUIdV4() (string, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_Z9SFmr28ft)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    uuid, err := random.UUIdV4()
    if err != nil {
        return
    }
    fmt.Println(uuid)
}
```

### <span id="RandUniqueIntSlice">RandUniqueIntSlice</span>

<p>生成一个不重复的长度为n的随机int切片。</p>

<b>函数签名:</b>

```go
func RandUniqueIntSlice(n, min, max int) []int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/uBkRSOz73Ec)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    result := random.RandUniqueIntSlice(5, 0, 10)
    fmt.Println(result) //[0 4 7 1 5] (random)
}
```

### <span id="RandFloat">RandFloat</span>

<p>生成随机float64数字，可以指定范围和精度。</p>

<b>函数签名:</b>

```go
func RandFloat(min, max float64, precision int) float64
```

<b>实例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zbD_tuobJtr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    floatNumber := random.RandFloat(1.0, 5.0, 2)
    fmt.Println(floatNumber) //2.14 (random number)
}
```

### <span id="RandFloats">RandFloats</span>

<p>生成随机float64数字切片，指定长度，范围和精度.</p>

<b>函数签名:</b>

```go
func RandFloats(n int, min, max float64, precision int) []float64
```

<b>实例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/I3yndUQ-rhh)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    floatNumbers := random.RandFloats(5, 1.0, 5.0, 2)
    fmt.Println(floatNumber) //[3.42 3.99 1.3 2.38 4.23] (random)
}
```