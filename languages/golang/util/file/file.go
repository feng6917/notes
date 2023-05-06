package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
  文件存在检测
  @author Bill
*/

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/*
*

	删除文件或文件夹
*/
func DeleteFile(abs_dir string) error {
	return os.RemoveAll(abs_dir)
}

func ReadFile(fp string) ([]byte, error) {
	var chunk []byte
	var n int
	//获得一个file
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println("read fail")
		return chunk, err
	}

	//把file读取到缓冲区中
	defer f.Close()
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err = f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return chunk, err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	return chunk, nil
	//fmt.Println(string(chunk))
}

func GenNewName(v, extName string) string {
	dirN, fileN := filepath.Split(v)
	return filepath.Join(dirN, fileN+extName)
}
