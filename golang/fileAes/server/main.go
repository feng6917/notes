package main

import (
	"flag"
	"fmt"
	"golang/fileAes/util"
	"golang/fileAes/util/file"
	"golang/fileAes/util/hex"
	"os"
	"strings"
)

const (
	extName   string = "-ccccccccccccpenny"
	saltValue string = "zxcavtbnwmlbkjhgfdesabnf"
	tip       string = "加密文件操作: \r\napp -p tip.txt -m t \r\n解密文件操作: \napp -p tip.txt -n n_tip.txt"
)

var (
	sourceFilePath string
	marshalFile    string
	deleteFile     string
	newFilePath    string
	defaultValue   bool
)

// 读取默认配置
func init() {
	flag.BoolVar(&defaultValue, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-", false, tip)
	flag.StringVar(&sourceFilePath, "p", "", "name: 源文件路径. \r\nvalue: 必传,不能为空.")
	flag.StringVar(&marshalFile, "m", "f", "name: 加密源文件. \r\nvalue: (t/f)非必传. ")
	flag.StringVar(&deleteFile, "d", "f", "name: 删除源文件. \r\nvalue: (t/f)非必传. ")
	flag.StringVar(&newFilePath, "n", "", "name: 新的文件路径. \r\nvalue: 非必传,默认为空.")
}

func app() {
	flag.Parse()
	var err error
	if strings.TrimSpace(sourceFilePath) == "" {
		fmt.Println(util.ErrorStr("文件名不能为空"))
		return
	}

	// 处理解密文件名格式
	if marshalFile == "f" {
		if !strings.HasSuffix(sourceFilePath, extName) {
			sourceFilePath += extName
		}
	}
	// 校验文件路径
	if !file.PathExists(sourceFilePath) {
		fmt.Println(util.ErrorStr("源文件不存在"))
		return
	}

	// 读取文件内容
	buf, err := file.ReadFile(sourceFilePath)
	if err != nil {
		fmt.Println(util.ErrorStr("读取文件内容失败"))
		return
	}
	var cbcBuf []byte
	var newFileName string
	if marshalFile == "t" {
		// 加密文件
		cbcBuf, err = hex.AesEncryptCBC(buf, []byte(saltValue))
		if err != nil {
			fmt.Println(util.ErrorStr("加密失败"))
			return
		}
		newFileName = file.GenNewName(sourceFilePath, extName)
	} else {
		cbcBuf, err = hex.AesDecryptCBC(buf, []byte(saltValue))
		if err != nil {
			fmt.Println(err)
			fmt.Println(util.ErrorStr("解密失败"))
			return
		}
		if newFilePath != "" {
			newFileName = newFilePath
		} else {
			newFileName = sourceFilePath[:len(sourceFilePath)-len(extName)]
		}
	}

	f, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(util.ErrorStr(err.Error()))
		return
	}
	_, err = f.Write(cbcBuf)
	if err != nil {
		fmt.Println(util.ErrorStr(err.Error()))
		return
	}
	fmt.Println("新文件创建成功。")
	if deleteFile == "t" {
		err = os.RemoveAll(sourceFilePath)
		if err != nil {
			fmt.Println(util.ErrorStr(err.Error()))
			return
		}
		fmt.Println("旧文件删除成功。")
	}
}
