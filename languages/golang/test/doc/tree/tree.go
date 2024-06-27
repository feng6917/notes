package main

import (
	"fmt"
	"sync"
)

/*
	今天跟同事聊起来后端返回前端树的问题，不考虑大数据量情况，一般数据量情况下如何更好的维护
	简单的写一个小例子进行测试，实际工作中需要针对数据进行分析，是否频繁操作，父子级哪个更适合用于key，是否针对一定量级特殊处理，分页等
	具体情况具体分析，没有什么是万能的解决方式，不断改进，不断优化。
*/

type Node struct {
	ID       uint32
	Name     string
	Level    int    // 层级
	ParentID uint32 // 父级ID
}

type CacheTree struct {
	Level sync.Map   // 层级映射 存储层级映射关系
	Data  []sync.Map // 数据映射 层级数据存储切片
}

// upInsertLevel 插入更新等级（同时创建空数据）
func (c *CacheTree) upInsertLevel(l int) {
	_, load := c.Level.LoadOrStore(l, l)
	if !load {
		l := len(c.Data)
		for i := l; i < l+1; i++ {
			c.Data = append(c.Data, sync.Map{})
		}
	}
}

// existLevel 是否存在该层级
func (c *CacheTree) existLevel(l int) bool {
	_, load := c.Level.LoadOrStore(l, l)
	return load
}

// storeData 存储数据
func (c *CacheTree) storeData(n Node) {
	c.Data[n.Level].Store(n.ID, n)
}

// UpInsert 更新创建
func (c *CacheTree) UpInsert(n Node) {
	// 层级存储
	c.upInsertLevel(n.Level)
	// 数据存储
	c.storeData(n)
}

// 数结构
type TreeNode struct {
	Node
	Child []TreeNode
}

// GetTree 取树
func (c *CacheTree) GetTree(level int, id uint32) TreeNode {
	var res TreeNode
	if level == 0 && id == 0 {
		// 根层级
		c.Data[0].Range(func(key, value interface{}) bool {
			k, kStatus := key.(uint32)
			if kStatus {
				res.Child = append(res.Child, c.GetTree(level, k))
			}
			return true
		})
	} else {
		// 指定层级
		c.Data[level].Range(func(key, value interface{}) bool {
			k, kStatus := key.(uint32)
			if kStatus && k == id {
				v, vStatus := value.(Node)
				if vStatus {
					res.Node = v
				}
			}
			return true
		})
		if c.existLevel(level + 1) {
			// 指定层级子层级
			c.Data[level+1].Range(func(key, value interface{}) bool {
				v, vStatus := value.(Node)
				if vStatus && v.ParentID == id {
					res.Child = append(res.Child, c.GetTree(level+1, v.ID))
				}
				return true
			})
		}
	}
	return res
}

func main() {

	n1 := Node{
		ID:       1,
		Name:     "1",
		Level:    0,
		ParentID: 0,
	}

	n2 := Node{
		ID:       2,
		Name:     "2",
		Level:    1,
		ParentID: 1,
	}

	n3 := Node{
		ID:       3,
		Name:     "3",
		Level:    2,
		ParentID: 2,
	}

	ct := CacheTree{}
	ct.UpInsert(n1)
	ct.UpInsert(n2)
	ct.UpInsert(n3)

	ts := ct.GetTree(0, 0)
	fmt.Println(ts)
}
