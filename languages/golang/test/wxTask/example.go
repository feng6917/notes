// 列表: 添加行, 删除选中行, 清空行, 排序, 表头表项文本居中
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"lgo/test/wxTask/task"

	"github.com/sirupsen/logrus"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a    *app.App
	w    *window.Window
	list *widget.List

	btn_normal  *widget.Button
	btn_desktop *widget.Button

	btn_add   *widget.Button
	btn_del   *widget.Button
	btn_clear *widget.Button

	btn_sendMsg *widget.Button

	edit_nameVal   *widget.Edit
	edit_textVal   *widget.Edit
	edit_sinceVal  *widget.Edit
	edit_numberVal *widget.Edit

	// dt_date *widget.DateTime
	// dt_time *widget.DateTime
)

var ta task.Task
var wx task.WX

type Info struct {
	RemarkName string
	Text       string
	Time       string
	Since      string
	Number     string
}

func main() {
	ta.Init()
	wx.Init()
	// Login

	// 导出日志
	{
		logName := fmt.Sprintf("wx_task_%s.log", time.Now().Format("20060102150405"))
		f, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return
		}
		defer func() {
			f.Close()
		}()
		// 编译exe后无法同时输出到控制台及文件
		// multiWriter := io.MultiWriter(os.Stdout, f)
		logrus.SetOutput(f)
	}

	go func() {
		tk := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-tk.C:
				ti, exist := ta.RangeTask()
				if exist && ti.RemarkName != "" {
					logrus.Infof("向 %s 发送 %s", ti.RemarkName, ti.Text)
					// 执行任务
					// go task.PlaySound(val)
					go wx.SendMsgByRemarkName(ti.RemarkName, ti.Text)
					ti.SendNumber += 1
					if ti.SendNumber == ti.Number {
						ta.Remove(ti.RemarkName)
					} else {
						ta.Ping(ti)
					}
					listRefresh()
				}

			}
		}
	}()

	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 1120, 400, "有封号风险,谨慎使用", 0, xcc.Window_Style_Default)
	// w.SetTextColor(xc.ABGR(220, 20, 60, 255))
	// xc.XWnd_SetCaptionTextColor(hWnd, 0xFF0000)
	w.SetTitleColor(xc.ABGR(220, 20, 60, 255))

	svgIcon := imagex.NewBySvgStringW(svgIcon)

	xc.XSvg_SetSize(svgIcon.GetSvg(), 23, 23)

	a.SetWindowIcon(svgIcon.Handle)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(2, 255, 8, false, 0)
	var startX int
	startX += 12

	btn_normal = widget.NewButton(int32(startX), 35, 55, 30, "网页登录", w.Handle)
	btn_normal.Event_BnClick1(onBnClick)
	btn_normal.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	btn_normal.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	startX += 60
	btn_desktop = widget.NewButton(int32(startX), 35, 65, 30, "客户端登录", w.Handle)
	btn_desktop.Event_BnClick1(onBnClick)
	btn_desktop.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	btn_desktop.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	startX += 70
	st := widget.NewShapeText(int32(startX), 40, 35, 30, "备注名", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 设置字体大小
	st.SetFont(font.New(10).Handle)

	// 创建名称编辑框
	startX += 45

	edit_nameVal = widget.NewEdit(startX, 35, 175, 30, w.Handle)
	edit_nameVal.SetTextColor(xc.ABGR(236, 64, 160, 255))
	edit_nameVal.SetToolTip("请输入朋友备注名称, 不能模糊检索")

	// 创建打开文件按钮
	startX += 185

	st = widget.NewShapeText(int32(startX), 40, 35, 30, "发送内容", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 设置字体大小
	st.SetFont(font.New(10).Handle)

	// 创建打开文件编辑框
	startX += 58

	edit_textVal = widget.NewEdit(startX, 35, 240, 30, w.Handle)
	edit_textVal.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit_textVal.SetToolTip("请输入发送消息内容")
	
	startX += 245
	
	st = widget.NewShapeText(int32(startX), 40, 35, 30, "发送间隔时间", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 创建间隔编辑框
	startX += 78

	edit_sinceVal = widget.NewEdit(startX, 35, 50, 30, w.Handle)
	edit_sinceVal.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit_sinceVal.SetToolTip("请输入间隔时间 单位: 秒, 不能小于10")

	// 创建次数按钮
	startX += 55
	
	st = widget.NewShapeText(int32(startX), 40, 35, 30, "发送次数", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 创建次数编辑框
	startX += 53

	edit_numberVal = widget.NewEdit(startX, 35, 50, 30, w.Handle)
	edit_numberVal.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit_numberVal.SetToolTip("请输入发送次数 , 不能小于0")

	// 创建List
	createList()

	startX += 60
	btn_add = widget.NewButton(int32(startX), 35, 55, 30, "添加任务", w.Handle)
	btn_add.Event_BnClick1(onBnClick)

	startX += 55 + 5
	btn_del = widget.NewButton(int32(startX), 35, 55, 30, "删除任务", w.Handle)
	btn_del.Event_BnClick1(onBnClick)

	startX += 55 + 5
	btn_clear = widget.NewButton(int32(startX), 35, 55, 30, "清空任务", w.Handle)
	btn_clear.Event_BnClick1(onBnClick)

	fileBuf, err := os.ReadFile("./init.json")
	if err != nil {
		logrus.Error("读取配置文件失败,err: ", err)
		panic("--------------------------")
	}

	var cf struct {
		Task []Info
	}
	err = json.Unmarshal(fileBuf, &cf)
	if err != nil {
		logrus.Error("解析配置文件失败,err: ", err)
		panic("--------------------------")
	}
	listAddItem(cf.Task)
	w.Show(true)
	a.Run()
	a.Exit()
}

// 按钮单击事件
func onBnClick(hEle int, pbHandled *bool) int {
	xc.XEle_Enable(hEle, false) // 操作前禁用按钮

	switch hEle {
	case btn_normal.Handle:
		err := wx.Login(false)
		if err != nil {
			a.MessageBox("登录", err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		}
	case btn_desktop.Handle:
		err := wx.Login(true)
		if err != nil {
			a.MessageBox("登录", err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		}
	case btn_add.Handle:
		AddItem()
	case btn_del.Handle:
		listDelSelectItem()
	case btn_clear.Handle:
		DeleteItem()
	}

	xc.XEle_Enable(hEle, true) // 操作后解禁按钮
	return 0
}

// 创建List
func createList() {
	// 创建List
	list = widget.NewList(10, 70, 1100, 315, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器: 5列
	list.CreateAdapter(9)
	// 列表_置项默认高度和选中时高度
	// list.SetItemHeightDefault(24, 26)
	list.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	list.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)
	// 列表_绘制项分割线
	// list.SetDrawItemBkFlags(xcc.List_DrawItemBk_Flag_Line | xcc.List_DrawItemBk_Flag_LineV | xcc.List_DrawItemBk_Flag_Leave | xcc.List_DrawItemBk_Flag_Stay | xcc.List_DrawItemBk_Flag_Select)
	// 表头和表项居中
	listTextAlign()

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY. 可参照例子: List2
	list.AddColumnText(50, "name1", "序号")
	list.AddColumnText(187, "name2", "备注名")
	list.AddColumnText(260, "name3", "发送内容")
	list.AddColumnText(127, "name4", "发送时间")
	list.AddColumnText(80, "name5", "发送间隔")
	list.AddColumnText(80, "name6", "发送次数")
	list.AddColumnText(80, "name7", "发送状态")
	list.AddColumnText(100, "name8", "待发送次数")
	list.AddColumnText(127, "name9", "创建时间")

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
					logrus.Info("列表当前排序: 倒序")
				} else {
					list.SetProperty("sortType", "1")
					logrus.Info("列表当前排序: 正序")
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
	err := wx.IsLogin()
	if err != nil {
		a.MessageBox("登录", "请登陆后进行操作", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	var nameVal, textVal, sinceVal, numberVal string
	vc := edit_nameVal.GetText(&nameVal, 100)
	if vc < 1 {
		a.MessageBox("提示", "备注名称不能小于5字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	if ta.Exist(nameVal) {
		a.MessageBox("提示", "该备注名称已存在任务中", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	if !wx.ExistRemarkName(nameVal) {
		a.MessageBox("提示", "未找到该备注名", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	textCount := edit_textVal.GetText(&textVal, 1000)
	if textCount < 3 {
		a.MessageBox("提示", "发送内容不能小于3字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	edit_sinceVal.GetText(&sinceVal, 1000)
	sinceInt, _ := strconv.Atoi(sinceVal)
	if sinceInt < 5 {
		a.MessageBox("提示", "间隔时间不能小于10秒", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	// if sinceInt > 360 {
	// 	a.MessageBox("提示", "间隔时间不能大于1小时", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
	// 	return
	// }

	edit_numberVal.GetText(&numberVal, 1000)
	numberInt, _ := strconv.Atoi(numberVal)
	if numberInt < 0 {
		a.MessageBox("提示", "发送次数不能小于1次", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	if numberInt > 1000 {
		a.MessageBox("提示", "发送次数不能大于1000次", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	// var yearInt, monthInt, dayInt int32
	// dt_date.GetDate(&yearInt, &monthInt, &dayInt)

	// var hourInt, minuteInt, secondInt int32
	// dt_time.GetTime(&hourInt, &minuteInt, &secondInt)

	// timeStart := time.Date(int(yearInt), time.Month(monthInt), int(dayInt), int(hourInt), int(minuteInt), int(secondInt), 0, time.Local)

	{
		edit_nameVal.SetText("")
		edit_nameVal.Redraw(true)

		edit_textVal.SetText("")
		edit_textVal.Redraw(true)

		// tn := time.Now()
		// dt_date.SetDate(int32(tn.Year()), int32(tn.Month()), int32(tn.Day()))
		// dt_time.SetTime(int32(tn.Hour()), int32(tn.Minute()), int32(tn.Second()))
		// dt_date.Redraw(true)
		// dt_time.Redraw(true)

		edit_sinceVal.SetText("")
		edit_sinceVal.Redraw(true)

		edit_numberVal.SetText("")
		edit_numberVal.Redraw(true)
	}

	num := list.GetCount_AD() + 1

	// 添加行
	var index int
	if list.GetProperty("sortType") == "1" { // 正序
		index = list.AddItemTextEx("name2", nameVal)
	} else { // 倒序
		index = list.InsertItemTextEx(0, "name2", nameVal)
	}
	logrus.Infof("添加行索引: %d", index)
	tn := time.Now()
	// 置行数据
	// 序号列设置int型的数据才能按数字大小排序
	list.SetItemInt(index, 0, num)
	list.SetItemText(index, 2, textVal)
	list.SetItemText(index, 3, fmt.Sprintf("%d/%d/%d %.2d:%.2d:%d", tn.Year(), tn.Month(), tn.Day(), tn.Hour(), tn.Minute(), tn.Second()))
	list.SetItemText(index, 4, sinceVal)
	list.SetItemText(index, 5, numberVal)
	// 发送状态
	// sendStatus := getSendStatus(timeStart, sinceInt, numberInt)
	sendStatus := "运行中"
	list.SetItemText(index, 6, sendStatus)
	// sn := getSendNumber(timeStart, sinceInt, numberInt)

	list.SetItemText(index, 7, strconv.Itoa(0))

	list.SetItemText(index, 8, time.Now().Format("2006/01/02 15:04:05"))

	list.Redraw(true)

	ti := task.TaskInfo{
		RemarkName: nameVal,
		Text:       textVal,
		Time:       tn,
		Since:      sinceInt,
		Number:     numberInt,
		SendStatus: sendStatus,
		SendNumber: 0,
	}
	ta.Ping(ti)
	ta.PrintTask()
	listRefresh()
}

func DeleteItem() {
	err := wx.IsLogin()
	if err != nil {
		a.MessageBox("登录", "请登陆后进行操作", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	list.DeleteItemAll()
	ta.Init()
	list.Redraw(true)
}

// List删除选中行
func listDelSelectItem() {
	err := wx.IsLogin()
	if err != nil {
		a.MessageBox("登录", "请登陆后进行操作", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

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
		val := list.GetItemText(int(indexArr[i]), 1)
		if val != "" {
			ta.Remove(val)
		}
		list.DeleteItem(int(indexArr[i]))
		logrus.Infof("删除行索引: %d\n", indexArr[i])

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
		iItem := int(indexArr[i])
		remarkName := list.GetItemText(iItem, 1)

		val, exist := ta.Live.Load(remarkName)
		if exist {
			ti, ok := val.(task.TaskInfo)
			if !ok {
				continue
			}
			list.SetItemText(int(indexArr[i]), 6, ti.SendStatus)
			list.SetItemText(int(indexArr[i]), 7, strconv.Itoa(ti.Number-ti.SendNumber))
		} else {
			list.SetItemText(int(indexArr[i]), 6, "已完成")
			list.SetItemText(int(indexArr[i]), 7, "0")
		}
		logrus.Infof("刷新行索引: %d 备注名: %s ", indexArr[i], remarkName)
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

func getListStatusWait(iItem int) (string, string) {
	ts := getListTime(iItem)
	since := list.GetItemText(iItem, 4)
	number := list.GetItemText(iItem, 5)
	sinceInt, numberInt := getStrInt(since, number)
	res1 := getSendStatus(*ts, sinceInt, numberInt)
	sn := getSendNumber(*ts, sinceInt, numberInt)
	return res1, strconv.Itoa(numberInt - sn)
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
			index = list.AddItemTextEx("name2", tasks[i].RemarkName)
		} else { // 倒序
			index = list.InsertItemTextEx(0, "name2", tasks[i].RemarkName)
		}
		logrus.Infof("添加行索引: %d", index)

		// 置行数据
		// 序号列设置int型的数据才能按数字大小排序
		list.SetItemInt(index, 0, num)
		list.SetItemText(index, 2, tasks[i].Text)
		t, ts := getInfoTime(tasks[i].Time)
		list.SetItemText(index, 3, ts)
		list.SetItemText(index, 4, tasks[i].Since)
		list.SetItemText(index, 5, tasks[i].Number)
		sinceInt, numInt := getStrInt(tasks[i].Since, tasks[i].Number)

		sendStatus := getSendStatus(t, sinceInt, numInt)
		list.SetItemText(index, 6, sendStatus)

		sn := getSendNumber(t, sinceInt, numInt)
		list.SetItemText(index, 7, strconv.Itoa(numInt-sn))
		ti := task.TaskInfo{
			RemarkName: tasks[i].RemarkName,
			Text:       tasks[i].Text,
			Time:       t,
			Since:      sinceInt,
			Number:     numInt,
			SendStatus: sendStatus,
			SendNumber: sn,
		}
		ta.Ping(ti)
		list.SetItemText(index, 8, tn.Format("2006/01/02 15:04:05"))
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

func getSendNumber(ts time.Time, since, num int) int {
	var res int
	tn := time.Now()
	for i := 1; i < num; i++ {
		if ts.Add(time.Second*time.Duration(i*since)).Sub(tn).Seconds() < 0 {
			res += 1
		} else {
			continue
		}
	}
	return res
}

func getSendStatus(ts time.Time, sinceInt, numberInt int) string {
	tn := time.Now()
	if ts.Add(time.Second*time.Duration(numberInt*sinceInt)).Sub(tn) > 0 {
		return "运行中"
	} else {
		return "已完成"
	}
}

func getStrInt(str1, str2 string) (int, int) {
	res1, _ := strconv.Atoi(str1)
	res2, _ := strconv.Atoi(str2)
	return res1, res2
}

const (
	svgStr1 = `<svg t="1715754216579" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5028" width="64" height="64"><path d="M134.8 757.6c16.5-167.3 115.5-240.1 209.8-268 9.4-2.8 11.8-15 4-21-69.2-53-90.4-151.7-52.5-233.1 5.4-11.7 12.1-22.7 20-32.9 6.4-8.2-0.2-20.1-10.6-19.2-44.8 3.8-87.1 31.4-108.7 77.8-30.2 64.7-14.2 142.9 39.4 186.6 7.4 6.1 5.2 17.9-4 20.8C149.5 494.4 64 563.2 64 725.6c0 24.9 15.3 36.9 57.2 42.8 0.5 0.1 1.1 0.1 1.6 0.1 6.2 0 11.3-4.7 12-10.9z m754.4 0c-16.5-167.3-115.5-240.1-209.8-268-9.4-2.8-11.8-15-4-21 69.2-53 90.4-151.7 52.5-233.1-5.4-11.7-12.1-22.7-20-32.9-6.4-8.2 0.2-20.1 10.6-19.2 44.8 3.8 87.1 31.4 108.7 77.8 30.2 64.7 14.2 142.9-39.4 186.6-7.4 6.1-5.2 17.9 4 20.8 82.7 25.8 168.2 94.6 168.2 257 0 24.9-15.3 36.9-57.2 42.8-0.5 0.1-1.1 0.1-1.6 0.1-6.2 0-11.3-4.7-12-10.9z" p-id="5029" fill="#1afa29"></path><path d="M655.5 414.3c-15.2 35.7-39.2 59.5-66.3 77.3 111.4 23.8 240.9 101.1 240.9 324.2 0 59.5-75.3 59.5-319.1 59.5-246.9 0-319.1 0-319.1-59.5 0-220.1 129.5-300.3 240.9-324.2C345.5 444 315.4 334 357.5 244.8s147.5-122 228.8-74.3c81.2 47.6 114.3 157.6 69.2 243.8z" p-id="5030" fill="#1afa29"></path></svg>`
	svgStr2 = `<svg t="1715754363569" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6184" width="64" height="64"><path d="M523.946667 85.333333C802.773333 85.333333 982.826667 307.541333 982.826667 511.338667c0 242.837333-197.397333 448-458.837334 448-84.16 0-150.464-16.597333-221.610666-55.402667l-88.618667 50.197333c-27.797333 8.426667-50.474667 5.162667-68.010667-9.834666-17.557333-14.997333-25.130667-34.730667-22.72-59.2a19570.688 19570.688 0 0 0 29.653334-106.368C123.008 741.76 64 656.426667 64 511.338667 64 307.541333 245.141333 85.333333 523.946667 85.333333z m-1.28 64C304.064 149.333333 128 317.12 128 522.666667c0 77.354667 24.874667 151.125333 70.634667 213.184l5.397333 7.125333 18.218667 23.509333-36.970667 128.746667a0.32 0.32 0 0 0 0.490667 0.384l113.408-63.829333 26.752 14.592C385.237333 878.72 452.544 896 522.666667 896 741.269333 896 917.333333 728.213333 917.333333 522.666667S741.269333 149.333333 522.666667 149.333333z m-192 320a53.333333 53.333333 0 1 1 0 106.666667 53.333333 53.333333 0 0 1 0-106.666667z m182.848 0a53.333333 53.333333 0 1 1 0 106.666667 53.333333 53.333333 0 0 1 0-106.666667z m182.869333 0a53.333333 53.333333 0 1 1 0 106.666667 53.333333 53.333333 0 0 1 0-106.666667z" fill="#1afa29" p-id="6185"></path></svg>`
	svgStr3 = `<svg t="1715755486537" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7317" width="64" height="64"><path d="M512 938.666667c235.637333 0 426.666667-191.029333 426.666667-426.666667S747.637333 85.333333 512 85.333333 85.333333 276.362667 85.333333 512s191.029333 426.666667 426.666667 426.666667zM329.376 649.376L480 498.741333V320a32 32 0 0 1 64 0v192a32 32 0 0 1-9.376 22.624l-160 160a32 32 0 1 1-45.248-45.248z" fill="#1afa29" p-id="7318"></path></svg>`
	svgStr4 = `<svg t="1715755995764" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11548" width="64" height="64"><path d="M512 288a32 32 0 0 1 31.488 26.24L544 320v136.576l-48 27.712-4.736 3.328a32 32 0 0 0-11.264 24.192V320a32 32 0 0 1 32-32z m326.464-211.2a547.008 547.008 0 0 1 108.8 108.8 32 32 0 1 1-51.2 38.4 483.008 483.008 0 0 0-96-96 32 32 0 0 1 38.4-51.2z" fill="#1296db" p-id="11549"></path><path d="M512 96a416 416 0 1 1 0 832 416 416 0 0 1 0-832z m0 64a352 352 0 1 0 0 704 352 352 0 0 0 0-704z m193.984 240a32 32 0 0 1-7.04 40.384l-4.672 3.328-166.272 96a32 32 0 0 1-36.736-52.096l4.736-3.328 166.272-96a32 32 0 0 1 43.712 11.712z" fill="#1296db" p-id="11550"></path></svg>`
	svgStr5 = `<svg t="1715755922301" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10415" width="64" height="64"><path d="M896.6 479.5c-17.8 0-29.6 11.8-29.6 29.6V589c0 100.6-79.9 183.4-183.4 183.4H225l38.5-35.5c5.9-5.9 11.8-14.8 11.8-23.7 0-17.8-14.8-32.5-32.5-32.5-8.9 0-20.7 3-26.6 8.9l-100.6 91.7c-14.8 11.8-11.8 32.5 0 44.4l100.6 88.8c5.9 5.9 17.8 11.8 26.6 11.8 17.8 0 32.5-11.8 35.5-29.6 0-11.8-5.9-20.7-14.8-26.6l-44.4-41.4h461.6c136.1 0 245.6-109.5 245.6-239.7v-80c-0.1-17.7-11.9-29.5-29.7-29.5z m-769.2 73.9c17.8 0 29.6-11.8 29.6-29.6V441c0-103.6 82.8-186.4 183.4-189.4H799l-38.5 35.5c-5.9 5.9-11.8 14.8-11.8 23.7 0 17.8 14.8 32.5 32.5 32.5 8.9 0 20.7-3 26.6-8.9l97.6-91.7c14.8-11.8 11.8-32.5 0-44.4l-100.6-88.8c-5.9-5.9-17.8-11.8-26.6-11.8-17.8 0-32.5 11.8-35.5 29.6 0 11.8 5.9 20.7 14.8 26.6l44.4 41.4H343.3c-136.1 0-245.6 109.5-245.6 248.5v79.9c0.1 17.9 11.9 29.7 29.7 29.7z m0 0" fill="#1296db" p-id="10416"></path><path d="M661.7 676.1h-71.6l-139-212.2c-7.2-11-12.4-19.9-15.7-26.7h-1.1c1.3 11.3 1.9 28.7 1.9 52.3v186.6H370V347.9h76.3l133.9 206.3c8.9 13.7 14.3 22.4 16.3 26.1h1.1c-1.4-7.9-2.1-23-2.1-45.3V348h66.3v328.1z" fill="#1296db" p-id="10417"></path></svg>`
	svgIcon = `<svg t="1715755531308" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8907" width="64" height="64"><path d="M135.1 961.2C60.6 961.2 0 900.6 0 826.1V396.3c0-74.5 60.6-135.1 135.1-135.1 15.5 0 30.8 2.6 45.3 7.8l-22.3 62.5c-7.3-2.6-15.1-3.9-23-3.9-37.9 0-68.7 30.8-68.7 68.7v429.9c0 37.9 30.8 68.7 68.7 68.7 37.9 0 68.7-30.8 68.7-68.7v-87.9h66.3v87.9c0.1 74.4-60.5 135-135 135zM887.5 961.2c-74.5 0-135.1-66.7-135.1-148.7v-77.2h66.3v77.2c0 45.4 30.8 82.4 68.7 82.4 37.9 0 68.7-36.9 68.7-82.4V409.9c0-45.4-30.8-82.3-68.7-82.3-8.8 0-17.3 2-25.3 5.8l-28.7-59.8c17.1-8.2 35.2-12.3 54-12.3 74.5 0 135.1 66.7 135.1 148.7v402.6c0.1 81.9-60.5 148.6-135 148.6z" fill="#ED4CA5" p-id="8908"></path><path d="M514.4 130.2c192 0 348.3 156.2 348.3 348.3 0 192-156.2 348.3-348.3 348.3S166.1 670.5 166.1 478.5c0-192.1 156.3-348.3 348.3-348.3m0-66.4C285.4 63.8 99.8 249.5 99.8 478.5s185.6 414.6 414.6 414.6c229 0 414.6-185.6 414.6-414.6 0-229-185.6-414.7-414.6-414.7z" fill="#ED4CA5" p-id="8909"></path><path d="M678.5 379.9c0-18.8 15.3-34.1 34.1-34.1 18.9 0 34.1 15.3 34.1 34.1v66c0 18.8-15.3 34.1-34.1 34.1s-34.1-15.3-34.1-34.1v-66zM282.1 379.9c0-18.8 15.3-34.1 34.1-34.1 18.9 0 34.1 15.3 34.1 34.1v66c0 18.8-15.3 34.1-34.1 34.1s-34.1-15.3-34.1-34.1v-66z m0 0M621.5 655c-43.5 0-78.9-35.4-78.9-78.9 0-3-0.6-5.7-1.6-8.4 16-9.3 26.9-26.4 26.9-46.2 0-29.6-24-53.5-53.5-53.5-29.6 0-53.5 24-53.5 53.5 0 19.8 10.9 36.9 26.9 46.2-1 2.6-1.6 5.4-1.6 8.4 0 43.5-35.4 78.9-78.9 78.9-13.1 0-23.7 10.6-23.7 23.7s10.6 23.7 23.7 23.7c45.1 0 84.7-23.9 107.1-59.6 22.4 35.7 61.9 59.6 107.1 59.6 13.1 0 23.7-10.6 23.7-23.7S634.6 655 621.5 655z" fill="#FFD500" p-id="8910"></path></svg>`
)
