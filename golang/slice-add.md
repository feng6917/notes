Slice1.7及以下扩容实操：

很早了解到Slice的扩容规则，根据内部元素长度来判断处理，具体如下：

1. 元素个数大等于1024，按照1.25，即 1/4 cap速率增长；
2. 元素个数小于1024，判断所需元素容量是否大于当前容量的2倍
   1. 大于，当前容量+所需容量
   2. 小于，当前容量的二倍

看起来很简单，理解起来很容易，但是实际上扩容真的仅仅只跟切片的容量有关么？

实操一下：

1. 首先初始化空的slice,类型是int

2. 添加2个元素，并打印容量

   ```
   空slice 新增2个元素后，长度为2，容量为2
   ```

3. 再次添加1个元素，并打印容量

   ```
   slice 再次新增1个元素后，新增元素小于当前容量*2，长度为3，容量为4
   ```

4. 以当前容量为原始容量进行下列操作

5. 原始容量添加3个元素，并打印容量

   ```
   原始容量新增3个元素，新增元素小于当前容量*2，长度为6，容量为8
   ```

6. 原始容量添加5个元素，并打印容量

   ```
   原始容量新增5个元素，新增元素+当前长度等于当前容量*2，长度为8，容量为8
   ```

7. 原始容量添加6个元素，并打印容量

   ```
   原始容量新增6个元素，新增元素+当前长度大于当前容量*2，长度为9，容量为10
   ```

8. 原始容量添加7个元素，并打印容量

   ```
   原始容量新增7个元素，新增元素+当前长度大于当前容量*2，长度为10，容量为10
   ```

9. 原始容量添加8个元素，并打印容量

   ```
   原始容量新增8个元素，新增元素+当前长度大于当前容量*2，长度为11，容量为12
   ```

10. 初始化容量大于1024的Slice并填充数据，新增若干元素并打印。

    ```
    大于1024元素，以1.25倍速率增长
    ```

    

```
package aaa

import "fmt"

func PrintSlice() {
   var slice []int
   slice = append(slice, 1, 2)
   fmt.Println("容量：", cap(slice))
   slice = append(slice, 3)

   fmt.Println("原容量：", cap(slice))

   is3 := []int{4, 5, 6}
   newSlice4 := AddElement(slice, is3...)
   fmt.Printf("新增: %d 容量: %d\n", len(is3), cap(newSlice4))

   is5 := []int{4, 5, 6, 7, 8}
   newSlice5 := AddElement(slice, is5...)
   fmt.Printf("新增: %d 容量: %d\n", len(is5), cap(newSlice5))

   is6 := []int{4, 5, 6, 7, 8, 9}
   newSlice6 := AddElement(slice, is6...)
   fmt.Printf("新增: %d 容量: %d\n", len(is6), cap(newSlice6))

   is7 := []int{4, 5, 6, 7, 8, 9, 10}
   newSlice7 := AddElement(slice, is7...)
   fmt.Printf("新增: %d 容量: %d\n", len(is7), cap(newSlice7))

   is8 := []int{4, 5, 6, 7, 8, 9, 10, 11}
   newSlice8 := AddElement(slice, is8...)
   fmt.Printf("新增: %d 容量: %d\n", len(is8), cap(newSlice8))

   bigIs := getBigInt()
   fmt.Println("大容量 原始: ", cap(bigIs))
   newBSlice8 := AddElement(bigIs, is8...)
   fmt.Printf("大容量 新增: %d 容量: %d\n", len(is8), cap(newBSlice8))
}

func AddElement(slice []int, e ...int) []int {
   return append(slice, e...)
}

func getBigInt() []int {
   res := make([]int, 0, 1026)
   for i := 0; i < 1025; i++ {
      res = append(res, i)
   }
   return res
}
```

总结：

- 大于等于1024元素扩容以1.25倍速率增长，并不是固定的1.25，部分可达1.561，之后慢慢速率将至1.249（1.25左右）
- 小于1024元素，且（新增元素数量+当前元素长度）<= 当前容量的二倍，扩容为当前容量二倍,并不是固定的2.0，部分可达2.365，之后慢慢速率将至2.0
- 小于1024元素，且 （新增元素数量+当前元素长度）> 当前容量的二倍
  - 当前长度+新增容量 为 奇数，扩容容量为 该奇数值 + 1  
  - 当前长度+新增容量 为 偶数，扩容后容量为 该偶数值  

扩容图：

![output](https://user-images.githubusercontent.com/82997695/222992202-2fdbd586-a5e4-4e2a-8c42-37b887e9648f.png)


1.20.1 扩容

1. 小于256个元素，跟1.17小于1024规则基本不变，增长单个元素部分可达2.359，之后慢慢将至2.0
2. 大于256个元素，采用平滑过度的扩容方式，避免出现1.17中，1024前扩容容量大于1024后扩容容量，部分可达2.365，之后慢慢下降，将至1.25
    扩容方式：
   ```
    const threshold = 256
		if oldCap < threshold {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < newLen {
				// Transition from growing 2x for small slices
				// to growing 1.25x for large slices. This formula
				// gives a smooth-ish transition between the two.
				newcap += (newcap + 3*threshold) / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = newLen
			}
		}
   ```    

测试扩容即折现图生成代码：
```
 package aaa

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"os"
)

func PrintSlice() {
	is8 := []int{4}
	var indexFs, SourceCs, NewCs []float64
	var minF, maxF float64
	var minf, maxf float64
	minf = 3
	maxf = 3
	var index int
	index = 1024
	for i := 1; i < 20000; i += 1 {
		bigInt := getBigIntNum(i)
		newBig := AddElement(bigInt, is8...)
		tempF := float64(cap(newBig)) / float64(cap(bigInt))
		fmt.Printf("------\n原始容量：%d，原始长度：%d 增加元素：%d，现在容量：%d, 增长率：%v \n",
			cap(bigInt), len(bigInt), len(is8), cap(newBig), tempF)
		indexFs = append(indexFs, float64(i))
		SourceCs = append(SourceCs, float64(cap(bigInt)))
		NewCs = append(NewCs, float64(cap(newBig)))
		if i < index {
			if minF < tempF {
				minF = tempF
			}
			if minf > tempF {
				minf = tempF
			}
		} else {
			if maxF < tempF {
				maxF = tempF
			}

			if maxf > tempF {
				maxf = tempF
			}
		}
	}
	fmt.Printf("小于 %d 最大增长速率:%f, 最小增长速率: %f\r\n", index, minF, minf)
	fmt.Printf("小于 %d 最大增长速率:%f, 最小增长速率: %f\r\n", index, maxF, maxf)

	ts1 := chart.ContinuousSeries{
		Name:    "Old Cap",
		XValues: indexFs,
		YValues: SourceCs,
	}

	ts2 := chart.ContinuousSeries{
		Name:    "New Cap",
		XValues: indexFs,
		YValues: NewCs,
	}
	genPng(ts1, ts2)
}

func AddElement(slice []int, e ...int) []int {
	return append(slice, e...)
}

func getBigIntNum(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}

func genPng(ts1, ts2 chart.ContinuousSeries) {
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 100,
			},
		},
		Series: []chart.Series{
			ts1,
			ts2,
		},
	}

	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	f, _ := os.Create("output.png")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	graph.Render(chart.PNG, f)
}

```

扩容图：


 
![output](https://user-images.githubusercontent.com/82997695/222992269-ceb7ecac-3a34-4043-ace2-2663532f9491.png)




