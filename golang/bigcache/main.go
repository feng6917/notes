package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/allegro/bigcache"
)

var (
	errUserNotInCache error = errors.New(" The user is not exist! ")
)

type user struct {
	ID        int64
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

type bigCache struct {
	users *bigcache.BigCache
}

func newBigCache() (*bigCache, error) {
	bCache, err := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 1 * time.Hour,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: false,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 256,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("new big cache: %w", err)
	}

	return &bigCache{
		users: bCache,
	}, nil
}

func (bc *bigCache) update(u user) error {
	bs, err := json.Marshal(&u)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return bc.users.Set(userKey(u.ID), bs)
}

func userKey(id int64) string {
	return strconv.FormatInt(id, 10)
}

func (bc *bigCache) read(id int64) (user, error) {
	bs, err := bc.users.Get(userKey(id))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return user{}, errUserNotInCache
		}

		return user{}, fmt.Errorf("get: %w", err)
	}

	var u user
	err = json.Unmarshal(bs, &u)
	if err != nil {
		return user{}, fmt.Errorf("unmarshal: %w", err)
	}

	return u, nil
}

func (bc *bigCache) delete(id int64) {
	bc.users.Delete(userKey(id))
}


/*
文章分析地址：
https://hackernoon.com/in-memory-caching-in-golang
code地址：
https://github.com/forPelevin/go-cache
Simple map ｜ gCache Library ｜ BigCache Library
*/
func main() {
	ce, err := newBigCache()
	if err != nil {
		panic(err)
	}
	u1 := user{
		ID:   1,
		Name: "huansijia",
	}
	err = ce.update(u1)
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	ur1, err := ce.read(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\r\n", ur1)

	ce.delete(1)

	_, err = ce.read(1)
	if err != nil {
		fmt.Println(err)
	}
}