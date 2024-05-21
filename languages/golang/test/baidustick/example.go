// 列表: 添加行, 删除选中行, 清空行, 排序, 表头表项文本居中
package main

import (
	"fmt"
	"os"
	"time"

	"lgo/test/baidustick/tb"

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
	a              *app.App
	w              *window.Window
	list           *widget.List
	btn_in         *widget.Button
	edit_cookieVal *widget.Edit
)
var cf struct {
	Cookie string
}

func main() {
	fname := "./cookie.txt"
	_, err := os.Stat(fname)
	if err != nil {
		of, _ := os.Create(fname)
		of.Close()
	} else {
		fileBuf, err := os.ReadFile(fname)
		if err != nil {
			logrus.Error("读取配置文件失败,err: ", err)
			panic("--------------------------")
		}

		cf.Cookie = string(fileBuf)
	}

	// 导出日志
	{
		logName := fmt.Sprintf("bt_login_%s.log", time.Now().Format("20060102150405"))
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
		tk := time.NewTicker(time.Hour * 12)
		for {
			select {
			case <-tk.C:
				logrus.Infof("开始签到")
				LoginIn()

			}
		}
	}()

	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 700, 400, "", 0, xcc.Window_Style_Default)
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

	st := widget.NewShapeText(int32(startX), 40, 35, 30, "Cookie", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 设置字体大小
	st.SetFont(font.New(10).Handle)

	// 创建名称编辑框
	startX += 55

	edit_cookieVal = widget.NewEdit(startX, 35, 555, 30, w.Handle)
	edit_cookieVal.SetTextColor(xc.ABGR(236, 64, 160, 255))
	edit_cookieVal.SetToolTip("请输入登录Cookie值, 不能为空")

	// 创建List
	createList()

	startX += 560
	btn_in = widget.NewButton(int32(startX), 35, 55, 30, "一键签到", w.Handle)
	btn_in.Event_BnClick1(onBnClick)
	btn_in.SetTextColor(xc.ABGR(236, 64, 160, 255))

	w.Show(true)
	a.Run()
	a.Exit()
}

// 按钮单击事件
func onBnClick(hEle int, pbHandled *bool) int {
	xc.XEle_Enable(hEle, false) // 操作前禁用按钮

	switch hEle {
	case btn_in.Handle:
		LoginIn()
	}

	xc.XEle_Enable(hEle, true) // 操作后解禁按钮
	return 0
}

// 创建List
func createList() {
	// 创建List
	list = widget.NewList(10, 70, 680, 315, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器: 5列
	list.CreateAdapter(6)
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
	list.AddColumnText(187, "name2", "吧名")
	list.AddColumnText(90, "name3", "经验值")
	list.AddColumnText(140, "name4", "等级")
	list.AddColumnText(76, "name5", "签到状态")
	list.AddColumnText(127, "name6", "签到时间")

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

// 一键签到
func LoginIn() {
	var cookieVal string
	cookieCount := edit_cookieVal.GetText(&cookieVal, 10000)
	if cookieCount < 10 {
		cookieVal = cf.Cookie
	}
	fmt.Println("======================================")
	logrus.Infof("cookie: %s", cookieVal)
	if len(cookieVal) < 10 {
		a.MessageBox("提示", "cookie不能小于10字符", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}

	res, err := tb.GetTb(cookieVal)
	if err != nil {
		logrus.Errorf("获取用户贴吧列表失败, error:%v", err)
		a.MessageBox("提示", err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		RewriteCookie("")
		return
	}

	RewriteCookie(cf.Cookie)

	logrus.Info("开始签到")
	// num := list.GetCount_AD() + 1
	logrus.Info("names count: ", len(res))
	if len(res) <= 0 {
		a.MessageBox("提示", "未获取到用户关注贴吧", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return
	}
	list.DeleteItemAll()
	{
		edit_cookieVal.SetText("")
		edit_cookieVal.Redraw(true)
	}
	for index, tmp := range res {
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("%d ----- %s 吧 开始签到\r\n", index, tmp)
		tbs, err := tb.GetTBSRequest(cf.Cookie)
		if err != nil {
			logrus.Errorf("获取用户贴吧TBS失败, error:%v", err)
			a.MessageBox("提示", "获取用户贴吧TBS失败, 请复制cookie到本地cookie.txt重试", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
			break
		}
		var hasSend bool
		err, hasSend = tb.SendPostFormFileRequest(cf.Cookie, tmp.Name, tbs)
		if err != nil {
			logrus.Errorf("%s 吧 签到失败, err: %v", tmp, err)
			continue
		}
		var status string
		var infoStr string
		if hasSend {
			status = "重复签到"
			infoStr = fmt.Sprintf("%d ----- %s 吧 亲，你之前已经签过了", index, tmp)
		} else {
			status = "签到成功"
			infoStr = fmt.Sprintf("%d ----- %s 吧 签到成功", index, tmp)
		}
		if infoStr != "" {
			logrus.Info(infoStr)
		}
		if list.GetProperty("sortType") == "1" { // 正序
			index = list.AddItemTextEx("name2", tmp.Name)
		} else { // 倒序
			index = list.InsertItemTextEx(0, "name2", tmp.Name)
		}
		logrus.Infof("添加行索引: %d", index)
		// 置行数据
		// 序号列设置int型的数据才能按数字大小排序
		list.SetItemInt(index, 0, index+1)
		fmt.Println(tmp)
		list.SetItemText(index, 2, tmp.CurExp)
		list.SetItemText(index, 3, fmt.Sprintf("%s-%s", tmp.BadgeTitle, tmp.BadgeLevel))
		list.SetItemText(index, 4, status)
		if status == "签到成功" {
			list.SetItemText(index, 5, time.Now().Format("2006/01/02 15:04:05"))
		}
		list.Redraw(true)
	}
}

func RewriteCookie(val string) {
	os.Remove("cookie.txt")
	of, _ := os.Create("cookie.txt")
	defer of.Close()
	of.WriteString(val)
}

const (
	svgIcon = `<svg t="1715920902253" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11776" width="64" height="64"><path d="M160.0512 566.1184h694.3232v281.7536H160.0512z" fill="#C4F000" p-id="11777"></path><path d="M728.4736 157.6448v-40.96h-81.92v41.216h-271.36v-40.96h-81.92v41.216H112.3328v733.9008H912.384V157.6448z m101.9904 652.4928H194.2528V239.5648h98.8672V286.72h81.92v-47.1552h271.36V286.72h81.92v-47.1552h102.144z" fill="#333333" p-id="11778"></path><path d="M471.808 591.5136L378.6752 482.0992l-62.4128 53.0944 150.6304 176.9472 249.5488-248.8832-57.856-58.0096-186.7776 186.2656z" fill="#333333" p-id="11779"></path></svg>`
)
