package main

import (
	"fmt"
	"sync"
)

/*
	今天跟同事聊起来后端返回前端树的问题，不考虑大数据量情况，一般数据量情况下如何更好的维护
	简单的写一个小例子进行测试，实际工作中需要针对数据进行分析，是否频繁操作，父子级哪个更适合用于key，是否针对一定量级特殊处理，分页等
	对于一个树需要考虑很多东西，广度，深度（大多数不需要考虑），假设场景，广度可以通过子级列表判断，深度可以通过level来判断
	考虑到节点数据检索，是否应该把该层级数据放到数据库中处理
	当一个层级数据几百以上时，就要首先考虑需求得合理性，用户能看这么多数据么，直接进行检索不香么
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

// autoRemoveNoParentIDNode
func (c *CacheTree) autoRemoveNoParentIDNode() {
	// 倒序处理
	for i := len(c.Data); i > 0; i-- {
		fmt.Println(i)
		c.removeCurrentNode(i)
	}
}

func (c *CacheTree) removeCurrentNode(l int) {
	if l-1 < 1 {
		return
	}
	// 获取父级ID列表
	m := make(map[uint32]struct{})
	c.Data[l-1].Range(func(key, value interface{}) bool {
		k, ok := key.(uint32)
		if ok {
			m[k] = struct{}{}
		}
		return true
	})

	var nodes []Node
	// 匹配当前级
	c.Data[l-1].Range(func(key, value interface{}) bool {
		val, ok := value.(Node)
		if ok {
			_, exist := m[val.ParentID]
			if !exist {
				nodes = append(nodes, val)
			}
		}
		return true
	})

	// 移除无关联数据
	c.removeNodes(nodes)

}

func (c *CacheTree) removeNodes(nodes []Node) {
	if len(nodes) == 0 {
		return
	}
	for _, node := range nodes {
		c.Data[node.Level].Delete(node.ID)
	}
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

// Remove 移除
func (c *CacheTree) Remove(l int, id uint32) {
	// 仅移除此节点 子级节点定时自动清理
	c.Data[l].Delete(id)
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

	// ts := ct.GetTree(0, 0)
	// fmt.Println(ts)

	for _, k := range ct.Data {
		k.Range(func(key, value interface{}) bool {
			fmt.Println("remove before: ", key, value)
			return true
		})
	}

	ct.Remove(1, 2)
	ct.autoRemoveNoParentIDNode()

	// ts1 := ct.GetTree(0, 0)
	// fmt.Println(ts1)
	for _, k := range ct.Data {
		k.Range(func(key, value interface{}) bool {
			fmt.Println("remove after: ", key, value)
			return true
		})
	}
}
