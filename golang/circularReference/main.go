package main

import (
	people2 "circularRe/Ans1/people"
	store2 "circularRe/Ans1/store"
	"circularRe/Ans2/other"
	"circularRe/Ans2/people"
	"circularRe/Ans2/store"
)

func main() {
	/*
		场景：客户购买商品，客户需要先看商品再拿钱，商家需要先拿钱再看商品，客户购买商品方法与商家售卖商品方法形成循环引用。
	*/
	//Source()

	// 新建公共接口包（父包），将需要循环调用的函数或方法抽象为接口
	//Q1()

	// 新建公共组合包（子包），在组合包中组合调用
	Q2()

}

//func Source() {
//	p := people.People{}
//	p.Buy()
//}

func Q1() {
	p := new(people2.People)
	s := new(store2.Store)
	p.StoreGoods = s
	s.PeopleMoney = p

	p.Buy()
	s.Sale()

}

func Q2() {
	p := new(people.People)
	s := new(store.Store)
	o := new(other.Other)
	o.PRepo = p
	o.SRepo = s
	o.Buy()

	o.Sale()

}
