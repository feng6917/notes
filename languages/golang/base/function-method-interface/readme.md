#### function-methond-interface

#### 目录

- <a href="#包初始化过程">包初始化过程</a>
- <a href="#函数方法接口之间关系">函数方法接口之间关系</a>
- <a href="#函数">函数</a>
- <a href="#方法">方法</a>
- <a href="#接口">接口</a>

#### 包初始化过程

```

    包初始化
    --------------------------------------
    Go语言程序的初始化和执行总是从 main.main 函数开始的。
    但是如果main包导入了其它的包，则会按照顺序将它们包 含进 main 包里(这里的导入顺序依赖具体实现，一般可能是以文件名或包路径名的字符串顺序导入).
    如果某个包被多 次导入的话，在执行的时候只会导入一次。

    当一个包被导入时，如果它还导入了其它的包，则先将其它的包包含进来，
    然后创建和初始化这个包的常量和变量,再调用包里的 init 函数，
    如果一个包有多个 init 函数的话，调用顺序未定义(实 现可能是以文件名的顺序调用)，同一个文件内的多个 init 则是以出现的顺序依次调用( init 不是普通函数，可以定 义有多个，所以也不能被其它函数调用)。
    最后，当 main 包的所有包级常量、变量被创建和初始化完成，并且 init 函数被执行后，才会进入 main.main 函数.

    import(other pkg (只导入一次)) const var init main

    注意！
    要注意的是，在 main.main 函数执行之前所有代码都运行在 同一个goroutine，也就是程序的主系统线程中。
    如果某 个 init 函数内部用go关键字启动了新的goroutine的话，新的 goroutine只有在进入 main.main 函数之后才可能被执行到。

```

#### 函数方法接口之间关系

```
    函数 方法 接口
    --------------------------------------
    Go语言中的函 数有具名和匿名之分:
    具名函数一般对应于包级的函数，是匿名函数的一种特例，当匿名函数引用了外部作用域中的变量时就成了闭包函数，闭包函数是函数式编程语言的核心。

    方法是 绑定到一个具体类型的特殊函数，Go语言中的方法是依托于 类型的，必须在编译时静态绑定。

    接口定义了方法的集合，这 些方法依托于运行时的接口对象，因此接口对应的方法是在运 行时动态绑定的。
    Go语言通过隐式接口机制实现了鸭子面向对象模型。
```

#### 函数

- 具名函数 Add()

  ```
    func Add(a, b int) int {
      return a + b
    }
  ```

- 匿名函数

  ```
    var Add = func(a, b int) int {
     return a + b
    }
    fmt.Println(Add(1, 2))
  ```

- 多个参数和多个返回值 Swap()

  ```
    func Swap(a, b int) (int, int) {
      return b, a
    }
  ```

- 可变数量的参数 Sum()

  ```
    // more 对应 []int 切片类型
    func Sum(a int, more ...int) int {
      for _, v := range more {
        a += v
      }
      return a
    }
  ```

- 给函数的返回值命名 Find()

  ```
    func Find(m map[int]int, key int) (value int, ok bool) {
      value, ok = m[key]
      return
    }
  ```

- 闭包的引用方式访问外部变量的行为可能会导致一些隐含的问题 A1() A2()

  ```
      func A1() {
        for i := 0; i < 3; i++ {
          i := i // 定义一个循环体内局部变量i
          defer func() {
            println(i)
          }()

        }
      }

      func A2() {
        for i := 0; i < 3; i++ { // 通过函数传入i
          // defer 语句会马上对调用参数求值
          defer func(i int) {
            println(i)
          }(i)

        }
      }
  ```

-

#### 方法

```
方法是由函数演变而来，只是将函数的第一个对象参数移动到了函数名前面了而已。将第一个函数参数移动到函数前面，从代码角度看虽然只是一 个小的改动，但是从编程哲学角度来看，Go语言已经是进入 面向对象语言的行列了。
对于给定的类型，每 个方法的名字必须是唯一的，同时方法和函数一样也不支持重载。

// 参数初始化
type repo struct {
 optionFirst string
}

// options options
type options func(*repo)

func NewRepo(opts ...options) *repo {
 srv := &repo{}
 for _, o := range opts {
  o(srv)
 }
 return srv
}

func withOptionFirst(t string) options {
 return func(c *repo) {
  c.optionFirst = t
 }
}
```

#### 接口

```
  Go的接口类型是对其它类型行为的抽象和概括；
  通过接口类型实现 了对鸭子类型的支持，使得安全动态的编程变得相对容易。
  Go语言中接口类型的独特之处在 于它是满足隐式实现的鸭子类型。
  鸭子类型说的是:只要走起路来像鸭子、叫起来也像鸭子，那么就可以把它当作鸭 子。
  Go语言中的面向对象就是如此，如果一个对象只要看起 来像是某种接口类型的实现，那么它就可以作为该接口类型使 用。
  这种设计可以让你创建一个新的接口类型满足已经存在的 具体类型却不用去破坏这些类型原有的定义；
```

```
    type UpperString string

    func (s UpperString) String() string {
      return strings.ToUpper(string(s))
    }
    有时候对象和接口之间太灵活了，导致我们需要人为地限制这种无意之间的适配。常见的做法是定义一个含特殊方法来区分接口。比如 runtime 包中的 Error 接口就定义了一个特有的
    RuntimeError 方法，用于避免其它类型无意中适配了该接口, 在protobuf中， Message 接口也采用了类似的方法，也定义了 一个特有的 ProtoMessage ，用于避免其它类型无意中适配了该接口(嵌套方式重用)
```

[Golang 目录](../../readme.md)
