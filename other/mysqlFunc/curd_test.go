package mysql_func

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Hobby    string    `json:"hobby"`
	Height   float64   `json:"height"`
	Age      int       `json:"age"`
	CreateAt time.Time `json:"create_at"`
}

func (c *User) TableName() string {
	return "user"
}

//func TestCreate(t *testing.T) {
//	u := User{
//		Name:     "001",
//		Hobby:    "h001",
//		Height:   16.8,
//		Age:      14,
//		CreateAt: time.Now(),
//	}
//
//	err := Create(&u, u.TableName())
//	fmt.Println("err: ", err)
//
//}
//
//func TestInterface2ArrayStr(t *testing.T) {
//	is := []int{1, 2, 3, 4, 5}
//	res, err := Interface2ArrayStr(is, true)
//	fmt.Printf("res: %v\r\n;err: %v\r\n", res, err)
//
//	ii := []string{"1", "22", "333", "4444"}
//	res, err = Interface2ArrayStr(ii, false)
//	fmt.Printf("res: %v\r\n;err: %v\r\n", res, err)
//}

type WhereReq struct {
	Id int `json:"id" in:"1" ft:"2"`
}

type UpdateReq struct {
	Name string `json:"name" ft:"1" in:"1"`
	Age  int    `json:"age" ft:"2" ig:"1" `
}

//func TestUpdate(t *testing.T) {
//	data := []interface{}{WhereReq{}, UpdateReq{}}
//	err := UnmarshalData(data)
//	if err != nil {
//		panic(err)
//	}
//	w := WhereReq{Id: 3}
//	u := UpdateReq{Name: "xiao a a"}
//	var user User
//	err = Update(w, u, user.TableName())
//	if err != nil {
//		fmt.Printf("err: %+v\r\n", err)
//	}
//}

func TestList(t *testing.T) {
	var err error
	data := []interface{}{WhereReq{}, UpdateReq{}}
	err = UnmarshalData(data)
	if err != nil {
		panic(err)
	}
	w := WhereReq{Id: 2}
	var user User
	var total int64
	var v []interface{}

	v, total, err = List(w, user.TableName(), nil)
	if err != nil {
		fmt.Printf("err: %+v\r\n", err)
	}
	fmt.Println(v)
	fmt.Println(total)
}
