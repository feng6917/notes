// 列表: 添加行, 删除选中行, 清空行, 排序, 表头表项文本居中
package main

import (
	"encoding/json"
	"fmt"
	"lgo/test/soundTask/task"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a    *app.App
	w    *window.Window
	list *widget.List

	btn_add   *widget.Button
	btn_del   *widget.Button
	btn_clear *widget.Button

	btn_openFile *widget.Button

	edit_nameVal     *widget.Edit
	edit_filenameVal *widget.Edit

	dt_date *widget.DateTime
	dt_time *widget.DateTime
)

var ta task.Task

type Info struct {
	Name     string
	Filename string
	Time     string
}

func main() {
	ta.Init()

	fileBuf, err := os.ReadFile("./init.json")
	if err != nil {
		fmt.Println("读取配置文件失败,err: ", err)
		panic("--------------------------")
	}

	var cf struct {
		Task []Info
	}
	err = json.Unmarshal(fileBuf, &cf)
	if err != nil {
		fmt.Println("解析配置文件失败,err: ", err)
		panic("--------------------------")
	}

	go func() {
		tk := time.NewTicker(time.Second * 3)
		for {
			select {
			case <-tk.C:
				key, val, exist := ta.RangeTask()
				if exist && key > 0 {
					logrus.Info(time.Unix(key, 0).Format("2006/01/02 15:04:05"))
					ta.Live.Delete(key)
					listRefresh()
					go task.PlaySound(val)

				}

			}
		}
	}()

	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 900, 400, "", 0, xcc.Window_Style_Default)
	svgIcon := imagex.NewBySvgStringW(svgIcon)
	xc.XSvg_SetSize(svgIcon.GetSvg(), 23, 23)
	a.SetWindowIcon(svgIcon.Handle)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(2, 255, 8, false, 0)
	var startX int
	startX += 15

	// 创建名称按钮
	nameButton := widget.NewButton(int32(startX), 35, 35, 30, "", w.Handle)

	svg1 := imagex.NewBySvgStringW(svgStr1)
	xc.XSvg_SetSize(svg1.GetSvg(), 28, 28)
	nameButton.SetIcon(svg1.Handle)
	nameButton.SetFocusBorderColor(0)
	nameButton.SetBorderSize(0, 0, 0, 0)
	nameButton.EnableDrawBorder(false)

	// 创建名称编辑框
	startX += 40

	edit_nameVal = widget.NewEdit(startX, 35, 130, 30, w.Handle)
	edit_nameVal.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit_nameVal.SetToolTip("请输入任务名称, 不能小于5字符且大于15字符")

	// 创建打开文件按钮
	startX += 135

	btn_openFile = widget.NewButton(int32(startX), 35, 35, 30, "", w.Handle)
	svg2 := imagex.NewBySvgStringW(svgStr2)
	xc.XSvg_SetSize(svg2.GetSvg(), 28, 28)
	btn_openFile.SetIcon(svg2.Handle)
	btn_openFile.Event_BnClick1(onBnClick)
	btn_openFile.EnableDrawBorder(false)

	// 创建打开文件编辑框
	startX += 40

	edit_filenameVal = widget.NewEdit(startX, 35, 240, 30, w.Handle)
	edit_filenameVal.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit_filenameVal.SetToolTip("请点击文件夹按钮进行播放文件设置")

	// 创建时间按钮
	startX += 245

	timeButton := widget.NewButton(int32(startX), 35, 35, 30, "", w.Handle)
	svg3 := imagex.NewBySvgStringW(svgStr3)
	xc.XSvg_SetSize(svg3.GetSvg(), 35, 35)
	timeButton.SetIcon(svg3.Handle)
	timeButton.SetFocusBorderColor(0)
	timeButton.SetBorderSize(0, 0, 0, 0)
	timeButton.EnableDrawBorder(false)

	// 创建日期 日期控件
	startX += 40

	dt_date = widget.NewDateTime(startX, 35, 90, 30, w.Handle)
	dt_date.SetStyle(0)

	// 创建日期 时间控件
	startX += 90

	dt_time = widget.NewDateTime(startX, 35, 90, 30, w.Handle)
	dt_time.SetStyle(1)

	// 创建List
	createList()

	startX += 100
	btn_add = widget.NewButton(int32(startX), 35, 55, 30, "添加任务", w.Handle)
	btn_add.Event_BnClick1(onBnClick)

	startX += 55 + 5
	btn_del = widget.NewButton(int32(startX), 35, 55, 30, "删除任务", w.Handle)
	btn_del.Event_BnClick1(onBnClick)

	startX += 55 + 5
	btn_clear = widget.NewButton(int32(startX), 35, 55, 30, "清空任务", w.Handle)
	btn_clear.Event_BnClick1(onBnClick)

	
	listAddItem(cf.Task)
	w.Show(true)
	a.Run()
	a.Exit()
}

// 按钮单击事件
func onBnClick(hEle int, pbHandled *bool) int {
	xc.XEle_Enable(hEle, false) // 操作前禁用按钮

	switch hEle {
	case btn_add.Handle:
		// listAddItem()
		AddItem()
	case btn_del.Handle:
		listDelSelectItem()
	case btn_clear.Handle:
		list.DeleteItemAll()
		ta.Init()
		list.Redraw(true)
	case btn_openFile.Handle:
		filename := wutil.OpenFile(w.Handle, []string{"*.mp3", "*.MP3", "*.wav", "*.WAV", "*.*"}, "")
		logrus.Infof("打开文件: %s", filename)
		edit_filenameVal.SetText(filename)
	}

	xc.XEle_Enable(hEle, true) // 操作后解禁按钮
	return 0
}

// 创建List
func createList() {
	// 创建List
	list = widget.NewList(10, 70, 880, 315, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器: 5列
	list.CreateAdapter(6)
	// 列表_置项默认高度和选中时高度
	list.SetItemHeightDefault(24, 26)
	// 列表_绘制项分割线
	// list.SetDrawItemBkFlags(xcc.List_DrawItemBk_Flag_Line | xcc.List_DrawItemBk_Flag_LineV | xcc.List_DrawItemBk_Flag_Leave | xcc.List_DrawItemBk_Flag_Stay | xcc.List_DrawItemBk_Flag_Select)
	// 表头和表项居中
	listTextAlign()

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY. 可参照例子: List2
	list.AddColumnText(50, "name1", "序号")
	list.AddColumnText(187, "name2", "名称")
	list.AddColumnText(330, "name3", "执行文件")
	list.AddColumnText(147, "name4", "执行时间")
	list.AddColumnText(67, "name5", "运行状态")
	list.AddColumnText(157, "name6", "创建时间")

	// 设置序号列可排序, 单击表头时排序
	list.SetSort(0, 0, true)
	// 这里我使用了置属性的方法是为了不新建多个变量, 因为考虑到组件可能会很多, 当然你也可以用变量来控制.
	// 这个置属性你可以理解为就是给元素绑定的map中赋值. 并不是在操作元素的属性.
	// 也是为了演示Set/GetProperty, 这个东西很有用, 比如说你的列表每1行都有隐藏的值, 就可以存在这里, 而不用自己新建一个map或slice, 看你需求了.
	list.SetProperty("sortType", "1") // 1是正序, 0是倒序.
	list.SetProperty("sortFlag", "0") // 只是我设定的标记

	// 列表头项被单击事件
	list.Event_LIST_HEADER_CLICK(func(iItem int32, pbHandled *bool) int {
		// 为了记录排序类型
		if iItem == 0 {
			// 下面这个sortFlag只是我设定的1个标记, 意义是让第1次单击表头排序时不设置sortType的值, 因为第1次默认就是正序
			if list.GetProperty("sortFlag") != "1" {
				list.SetProperty("sortFlag", "1")
			} else {
				if list.GetProperty("sortType") == "1" {
					list.SetProperty("sortType", "0")
					fmt.Println("列表当前排序: 倒序")
				} else {
					list.SetProperty("sortType", "1")
					fmt.Println("列表当前排序: 正序")
				}
			}
		}
		return 0
	})
}

// 表头和表项居中, 纯代码实现需要记一些api, 需要有清晰的思维, 还是用设计器来的简单, 真要写大程序不可能离开设计器的
func listTextAlign() {
	list.Event_LIST_HEADER_TEMP_CREATE_END(func(pItem *xc.List_Header_Item_, pbHandled *bool) int {
		for i := 0; i < list.GetColumnCount(); i++ {
			hEle := list.GetHeaderTemplateObject(i, 1)
			if a.IsHXCGUI(hEle, xcc.XC_SHAPE_TEXT) { // 是形状文本
				xc.XShapeText_SetTextAlign(hEle, xcc.TextAlignFlag_Center|xcc.TextAlignFlag_Vcenter)
			}
		}
		return 0
	})

	list.Event_LIST_TEMP_CREATE_END(func(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
		// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		if nFlag == 1 {
			for i := 0; i < list.GetColumnCount(); i++ {
				hEle := list.GetTemplateObject(int(pItem.Index), i, 1)
				if a.IsHXCGUI(hEle, xcc.XC_SHAPE_TEXT) { // 是形状文本
					xc.XShapeText_SetTextAlign(hEle, xcc.TextAlignFlag_Center|xcc.TextAlignFlag_Vcenter)
				}
			}
		}
		return 0
	})
}

// 添加数据
func AddItem() {
	var nameVal, filenameVal string
	vc := edit_nameVal.GetText(&nameVal, 100)
	if vc < 5 {
		a.MessageBox("提示", "任务名称不能小于5字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	if vc > 15 {
		a.MessageBox("提示", "任务名称不能大于15字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	filenameCount := edit_filenameVal.GetText(&filenameVal, 1000)
	if filenameCount < 3 {
		a.MessageBox("提示", "播放文件路径不能小于3字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	lowVal := strings.ToLower(filenameVal)
	if !strings.HasSuffix(lowVal, ".mp3") && !strings.HasSuffix(lowVal, ".wav") {
		a.MessageBox("提示", "播放文件路径不符合格式, 必须为.mp3或者.wav", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	var yearInt, monthInt, dayInt int32
	dt_date.GetDate(&yearInt, &monthInt, &dayInt)

	var hourInt, minuteInt, secondInt int32
	dt_time.GetTime(&hourInt, &minuteInt, &secondInt)

	cd := time.Date(int(yearInt), time.Month(monthInt), int(dayInt), int(hourInt), int(minuteInt), int(secondInt), 0, time.Local)
	if cd.Sub(time.Now()).Minutes() < 1 {
		a.MessageBox("提示", "任务间隔时间小于1分钟,请设置大于1分钟的任务间隔时间", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	if ta.Exist(cd.Unix()) {
		a.MessageBox("提示", "该执行时间已经设置任务,请重新设置任务执行时间", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	{
		edit_nameVal.SetText("")
		edit_nameVal.Redraw(true)
		edit_filenameVal.SetText("")
		edit_filenameVal.Redraw(true)
		tn := time.Now()
		dt_date.SetDate(int32(tn.Year()), int32(tn.Month()), int32(tn.Day()))
		dt_time.SetTime(int32(tn.Hour()), int32(tn.Minute()), int32(tn.Second()))
		dt_date.Redraw(true)
		dt_time.Redraw(true)
	}

	num := list.GetCount_AD() + 1

	// 添加行
	var index int
	if list.GetProperty("sortType") == "1" { // 正序
		index = list.AddItemTextEx("name2", nameVal)
	} else { // 倒序
		index = list.InsertItemTextEx(0, "name2", nameVal)
	}
	fmt.Printf("添加行索引: %d\n", index)

	// 置行数据
	// 序号列设置int型的数据才能按数字大小排序
	list.SetItemInt(index, 0, num)
	list.SetItemText(index, 2, filenameVal)
	list.SetItemText(index, 3, fmt.Sprintf("%d/%d/%d %.2d:%.2d:%d", yearInt, monthInt, dayInt, hourInt, minuteInt, secondInt))

	list.SetItemText(index, 4, "运行中")
	list.SetItemText(index, 5, time.Now().Format("2006/01/02 15:04:05"))

	list.Redraw(true)

	ta.Ping(cd.Unix(), filenameVal)

	listRefresh()
}

// List删除选中行
func listDelSelectItem() {
	count := list.GetSelectItemCount()
	if count == 0 {
		w.MessageBox("提示", "你没有选中列表任何行!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
		return
	}

	var indexArr []int32
	// 取选中行索引数组
	list.GetSelectAll(&indexArr, count)
	// 根据选中行索引数组倒着删, 正着删的话你每删1行下面的行索引就变了
	for i := count - 1; i > -1; i-- {
		vt := getListTime(int(indexArr[i]))
		if vt != nil {
			ta.Live.Delete(vt.Unix())
		}
		list.DeleteItem(int(indexArr[i]))
		fmt.Printf("删除行索引: %d\n", indexArr[i])

	}

	// 重排剩余行序号
	count = list.GetCount_AD()
	if list.GetProperty("sortType") == "1" { // 正序
		for i := 0; i < count; i++ {
			list.SetItemInt(i, 0, i+1)
		}
	} else { // 倒序
		for i, num := 0, count; i < count; i, num = i+1, num-1 {
			list.SetItemInt(i, 0, num)
		}
	}

	// 刷新列表项数据
	list.RefreshData()
	// 列表重绘
	list.Redraw(true)
}

func listRefresh() {
	list.SetSelectAll()
	count := list.GetSelectItemCount()
	if count == 0 {
		w.MessageBox("提示", "你没有选中列表任何行!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
		return
	}

	var indexArr []int32
	// 取选中行索引数组
	list.GetSelectAll(&indexArr, count)
	// 根据选中行索引数组倒着删, 正着删的话你每删1行下面的行索引就变了
	for i := count - 1; i > -1; i-- {
		vt := getListTime(int(indexArr[i]))
		if vt == nil {
			continue
		}
		if vt.Sub(time.Now()) < 0 {
			list.SetItemText(int(indexArr[i]), 4, "已过期")
		}
		fmt.Printf("刷新行索引: %d\n", indexArr[i])
	}

	list.CancelSelectAll()

	// 刷新列表项数据
	list.RefreshData()
	// 列表重绘
	list.Redraw(true)
}

func getListTime(iItem int) *time.Time {
	var res time.Time
	val := list.GetItemText(iItem, 3)
	l1 := strings.Split(val, " ")
	if len(l1) != 2 {
		return nil
	}
	d1 := strings.Split(l1[0], "/")
	if len(d1) != 3 {
		return nil
	}
	t1 := strings.Split(l1[1], ":")
	if len(t1) != 3 {
		return nil
	}
	yearInt, _ := strconv.Atoi(d1[0])
	monthInt, _ := strconv.Atoi(d1[1])
	dayInt, _ := strconv.Atoi(d1[2])
	hourInt, _ := strconv.Atoi(t1[0])
	minuteInt, _ := strconv.Atoi(t1[1])
	secondInt, _ := strconv.Atoi(t1[2])

	res = time.Date(yearInt, time.Month(monthInt), dayInt, hourInt, minuteInt, secondInt, 0, time.Local)
	return &res
}

func listAddItem(tasks []Info) {
	// 循环添加数据
	count := len(tasks)
	tn := time.Now()
	for i := 0; i < count; i++ {
		num := list.GetCount_AD() + 1

		// 添加行
		var index int
		if list.GetProperty("sortType") == "1" { // 正序
			index = list.AddItemTextEx("name2", tasks[i].Name)
		} else { // 倒序
			index = list.InsertItemTextEx(0, "name2", tasks[i].Name)
		}
		fmt.Printf("添加行索引: %d\n", index)

		// 置行数据
		// 序号列设置int型的数据才能按数字大小排序
		list.SetItemInt(index, 0, num)
		list.SetItemText(index, 2, tasks[i].Filename)
		t, ts := getInfoTime(tasks[i].Time)
		list.SetItemText(index, 3, ts)
		if t.Sub(tn) > 0{
			list.SetItemText(index, 4, "运行中")
			ta.Ping(t.Unix(), tasks[i].Filename)
		}else{
			list.SetItemText(index, 4, "已过期")
		}
		list.SetItemText(index, 5, tn.Format("2006/01/02 15:04:05"))
	}

	list.Redraw(true)
}

func getInfoTime(str string) (time.Time, string) {
	var res string
	var resTime time.Time
	l := strings.Split(str, ":")
	if len(l) != 2 {
		return resTime, res
	}
	hourInt, _ := strconv.Atoi(l[0])
	minuteInt, _ := strconv.Atoi(l[1])

	tn := time.Now()
	return time.Date(tn.Year(), tn.Month(), tn.Day(), hourInt, minuteInt, 0, 0, tn.Location()), fmt.Sprintf("%d/%d/%d %.2d:%.2d:00", tn.Year(), tn.Month(), tn.Day(), hourInt, minuteInt)
}

const (
	svgStr1 = `<svg t="1715582664333" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4681" xmlns:xlink="http://www.w3.org/1999/xlink" width="64" height="64"><path d="M512 1024C229.236 1024 0 794.764 0 512S229.236 0 512 0s512 229.236 512 512-229.236 512-512 512z m0-971.636C258.327 52.364 52.364 258.327 52.364 512S258.327 971.636 512 971.636 971.636 765.673 971.636 512 765.673 52.364 512 52.364z m229.236 670.254c-6.981 0-13.963-2.327-18.618-8.145l-84.945-154.764H387.49l-84.946 154.764c-5.818 5.818-11.636 8.145-18.618 8.145-6.982 0-13.963-2.327-18.618-8.145-10.473-10.473-10.473-26.764 0-37.237l82.618-148.945c1.164-4.655 3.491-10.473 8.146-13.964l132.654-239.709c1.164-3.49 3.491-8.145 6.982-11.636 4.655-4.655 11.636-8.146 18.618-8.146 6.982 0 13.964 2.328 18.618 8.146 3.491 3.49 5.819 8.145 6.982 12.8L762.182 678.4c4.654 4.655 6.982 10.473 6.982 17.455s-2.328 13.963-8.146 18.618c-5.818 5.818-12.8 8.145-19.782 8.145zM608.582 509.673L512 335.127l-96.582 174.546h193.164z" p-id="4682" fill="#f4ea2a"></path></svg>`
	svgStr2 = `<svg t="1715583937761" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5714" width="64" height="64"><path d="M928.426667 458.453333a61.013333 61.013333 0 0 0-51.626667-27.733333h-59.52v-58.666667a100.053333 100.053333 0 0 0-98.346667-101.333333H498.133333L344.106667 170.666667H183.68A100.053333 100.053333 0 0 0 85.333333 272v539.946667A42.666667 42.666667 0 0 0 128 853.333333h620.8a62.08 62.08 0 0 0 56.96-37.333333l128-300.16a59.946667 59.946667 0 0 0-5.333333-57.386667zM128 713.813333V272A57.173333 57.173333 0 0 1 183.68 213.333333h147.84l154.026667 100.053334H718.933333a57.173333 57.173333 0 0 1 55.68 58.666666v58.666667H277.333333a62.08 62.08 0 0 0-58.026666 39.893333zM894.506667 499.2l-77.226667 180.906667-50.773333 119.04a19.2 19.2 0 0 1-17.706667 11.52l-612.053333 1.92 121.813333-326.4a19.413333 19.413333 0 0 1 18.133333-12.16H876.8a18.986667 18.986667 0 0 1 16.213333 8.533333 18.346667 18.346667 0 0 1 1.493334 16.64z" p-id="5715" fill="#f4ea2a"></path></svg>`
	svgStr3 = `<svg t="1715656463211" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2386" width="64" height="64"><path d="M554.666667 537.6l115.2 115.2-29.866667 29.866667-128-128V341.333333h42.666667v196.266667z m-21.333334 358.4C332.8 896 170.666667 733.866667 170.666667 533.333333S332.8 170.666667 533.333333 170.666667 896 332.8 896 533.333333 733.866667 896 533.333333 896z m0-42.666667c174.933333 0 320-145.066667 320-320S708.266667 213.333333 533.333333 213.333333 213.333333 358.4 213.333333 533.333333 358.4 853.333333 533.333333 853.333333z" fill="#f4ea2a" p-id="2387"></path></svg>`
	svgIcon = `<svg t="1715737056018" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8692" width="64" height="64"><path d="M128 0h64v128a64 64 0 0 0 64 64h512a64 64 0 0 0 64-64V0h64a128 128 0 0 1 128 128v768a128 128 0 0 1-128 128H128a128 128 0 0 1-128-128V128a128 128 0 0 1 128-128z m641.152 355.776l-303.936 303.936-175.872-205.184-72.896 62.464 243.328 283.904 377.216-377.216-67.84-67.904zM768 0v64a64 64 0 0 1-64 64H320a64 64 0 0 1-64-64V0h512z" fill="#66C23A" p-id="8693"></path></svg>`
)
