package main

import (
	"flag"
	"fmt"
	"lgo/base/gorm_server"
	"lgo/base/gorm_server/config"
	"lgo/base/gorm_server/service"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var configPath string

// 读取默认配置
func init() {
	flag.StringVar(&configPath, "-f", "./app.conf", "配置路径,为空默认")
}

func main() {
	var conf *config.Config
	var err error
	conf, err = config.Init(configPath)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\r\n", conf)
	}

	var m *gorm_server.Manager
	dc := conf.DBConfigs[conf.EnvConfig.Mode]
	m, err = gorm_server.NewManager(dc)
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	tempPrefix := r.PathPrefix("/api").Subrouter()
	people := tempPrefix.PathPrefix("/people").Subrouter()
	people.Path("/create").Methods("POST").Handler(service.CreateData(m))
	people.Path("/update").Methods("PUT").Handler(service.UpdatePeople(m))
	people.Path("/delete").Methods("DELETE").Handler(service.DeletePeople(m))

	err = http.ListenAndServe(fmt.Sprintf(":%d", conf.EnvConfig.Port), r)
	if err != nil {
		panic(err)
	}
	//
	////  新增数据测试
	//p := model.People{
	//	Name:      "test001",
	//	Password:  "pwd001",
	//}
	//m.PeopleRepo.CreateData(&p)

	// 增
	//data.Create(&Product{Code: "L1212", Price: 1000})

	// 查
	//var product Product
	//data.First(&product, 1) // 找到id为1的产品
	//data.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品
	//
	//// 改 - 更新产品的价格为 2000
	//data.Model(&product).Update("Price", 2000)
	//
	//// 删 - 删除产品
	//data.Delete(&product)
}
