

##### 测试，了解一下

​	这里说的测试是指软件测试，其目的在于检验软件是否满足或弄清预期结果与实际结果之间的差别。即为了发现程序中的错误而使用热工或自动化手段执行程序或测试某个系统的过程。

   测试整个过程太过于繁杂，我们可以进行拆解，分为一个个阶段，一个个单元，一条条示例，对于单个那就容易太多了。大概处理逻辑示例搭建一座房子，把房子分为装修，水工，木工，电工等，类比一个个阶段，厨房，客厅，卧室，卫生间等，类比一个个单元，地板，墙壁，天花板等，类比一条条示例，由小及大，由少及多，一个整个流程下来其实也不是那么复杂。

  测试有没有必要，答案是肯定的，理由数不胜数。稳定性，快速迭代，重构，维护成本等，不作具体说明。

##### 测试，尝试一下

​	多种测试类型，参考[Go单测从零到溜系列0—单元测试基础](https://www.liwenzhou.com/posts/Go/unit-test-0/), 杜绝cv，引入之前写过的一句话，你现在或者将来要写的代码可能有前辈已经写过了一万零一遍。

##### 测试，参考一下

​	测试会写了，但如何在项目中更好的使用呢，推荐一下我一直使用的测试写法，分两步，1. 根据不同模块，构建基础模块测试初始化。 2. 分别创建不同模块测试文件，编写测试。

   示例：


- db.go
```
package main

    import (
        "time"

        log "github.com/sirupsen/logrus"
        "gorm.io/gorm"
    )

    type Db struct {
        db *gorm.DB
    }

    func (c *Db) Init() {
        log.Info("数据层初始化成功")
    }

    func (c *Db) Insert() {
        log.Infof("数据层 %v 插入一条数据完成", time.Now().Format("2006/01/02 15:04:05"))
    }
```

- db_test.go
```
package main

    import "testing"

    func TestInsert(t *testing.T) {
        db.Insert()
    }
```

- service.go
```
package main

    import (
        "time"

        log "github.com/sirupsen/logrus"
    )

    type Service struct {
        db *Db
    }

    func (c *Service) Init(db *Db) {
        c.db = db
        log.Info("服务层初始化成功")
    }
    func (c *Service) CreateUser() {
        c.db.Insert()
        log.Infof("服务层 %v 创建一个用户", time.Now().Format("2006/01/02 15:04:05"))
    }
```

- service_test.go
```
package main

    import "testing"

    func TestCreateUser(t *testing.T) {
        s.CreateUser()
    }
```

- example.go
```
package main

    func main() {
        // db := Db{}
        // db.Init()
        // s := Service{}
        // s.Init(&db)
        // m := Manager{}
        // m.Init(&s)

    }
```

- example_test.go

```
	package main

    import (
        "fmt"
        "os"
        "testing"
    )

    var m *Manager

    var s *Service

    var db *Db

    const (
        testMode string = "service"
    )

    func setup() {
        switch testMode {
        case "db":
            db = &Db{}
            db.Init()
        case "service":
            db = &Db{}
            db.Init()
            s = &Service{}
            s.Init(db)
        case "manager":
            // db = &Db{}
            // db.Init()
            // s = &Service{}
            // s.Init(db)
            // m = &Manager{}
            // m.Init(s)
        case "external_grpc":

        case "internal_grpc":

        case "http":
            fmt.Println("--------- http --------------")
        }

        fmt.Println("Before all tests")
    }

    func teardown() {
        switch testMode {
        case "grpc":
            // testConn.Close()
        }
        fmt.Println("After all tests")
    }

    func TestMain(m *testing.M) {
        setup()
        code := m.Run()
        teardown()
        os.Exit(code)
    }

```



​	
