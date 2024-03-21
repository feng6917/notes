package bigcache

import (
	"testing"
)

/*
文章分析地址：
https://hackernoon.com/in-memory-caching-in-golang
code地址：
https://github.com/forPelevin/go-cache
Simple map ｜ gCache Library ｜ BigCache Library
*/

func TestBigCache_Read(t *testing.T) {
	bc, err := newBigCache()
	if err != nil {
		t.Error(err.Error())
	}
	u := user{
		ID:   1,
		Name: "bigCache",
	}
	err = bc.Update(u)
	if err != nil {
		t.Error(err)
	}

	bu, err := bc.Read(1)
	if err != nil {
		t.Error(err)
	}

	if u.ID != bu.ID || u.Name != bu.Name {
		t.Errorf("got %+v, want:%+v", bu, u)
	}

}

func TestBigCache_Delete(t *testing.T) {
	bc, err := newBigCache()
	if err != nil {
		t.Error(err.Error())
	}
	u := user{
		ID:   1,
		Name: "bigCache",
	}
	err = bc.Update(u)
	if err != nil {
		t.Error(err)
	}

	bc.Delete(u.ID)

	bu, err := bc.Read(1)
	if err == nil || bu.ID > 0 {
		t.Errorf("got %+v, want:%+v", bu.ID, 0)
	}
}
