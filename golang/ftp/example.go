package ftp

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/secsy/goftp"
)

type SourceMonitor struct {
	Name       string   `json:"name"`
	ImagePaths []string `json:"image_paths"`
}

func example() {
	serverAddr := "10.0.0.7"
	user := "ftpzhst"
	pwd := "ftp1111"
	cli, err := newClient(serverAddr, user, pwd)
	if err != nil {
		fmt.Println("init client err: ", err)
		panic("------ init fail ----")
	}
	fmt.Println("init err: ", err)

	ns := "/monitor" // 指定文件夹
	// 上传
	Upload(cli, ns)
	//cli.Readdir(ns)

	//os.Open("./source.json")

	//cli.FtpDel(path.Join(ns, "go.mod"))

}

func Upload(cli *ftpClient, ns string) {
	lf := "./go.mod"
	fs, err := cli.FtpUpload(ns, lf)
	if err != nil {
		fmt.Println("ftp upload err: ", err)
		panic("------------ upload fail -------")
	}
	fmt.Println("upload err: ", err)
	fmt.Println("fs: ", fs)
}

type ftpClient struct {
	*goftp.Client
}

func newClient(serverAddr, user, password string) (*ftpClient, error) {
	ftp, err := goftp.DialConfig(goftp.Config{
		User:               user,
		Password:           password,
		ConnectionsPerHost: 10,
		Timeout:            time.Second * 30,
	}, serverAddr)
	if err != nil {
		return nil, errors.New("Fail to conn to ftp server " + serverAddr + " due to " + err.Error())
	}
	return &ftpClient{Client: ftp}, nil
}

func (cli ftpClient) FtpUpload(ns, localFile string) (string, error) {
	buf, err := os.ReadFile(localFile)
	if err != nil {
		return "", err
	}
	fmt.Println("up 1")
	savePath, err := cli.Mkdir(ns)
	if err != nil {
		return "", err
	}
	fmt.Println("up 2")
	fmt.Println("savePath: ", savePath)

	// 上传完毕后关闭当前的ftp连接
	defer cli.Client.Close()
	_, fname := filepath.Split(localFile)
	dstPath := path.Join(savePath, fname)
	// 文件上传
	fmt.Println(dstPath)
	return dstPath, cli.Client.Store(dstPath, bytes.NewReader(buf))
}

func (cli ftpClient) GetPath(ns string) (string, error) {
	// 2.得到pwd 当前路径
	pwd, pwdErr := cli.Client.Getwd()
	if pwdErr != nil {
		return "", pwdErr
	}

	// 2. 创建savePath
	savePath := path.Join(pwd, ns)
	return savePath, nil
}

func (cli ftpClient) Mkdir(ns string) (string, error) {
	// 2.得到pwd 当前路径
	savePath, getPathErr := cli.GetPath(ns)
	if getPathErr != nil {
		return "", getPathErr
	}

	_, err := cli.Client.Mkdir(savePath)
	if err != nil {
		fmt.Println("mkdir err: ", err)
		// 由于搭建ftp的时候已经给了`pwd` 777的权限，这里忽略文件夹创建的错误
		if !strings.Contains(err.Error(), "550-Create directory operation failed") {
			return savePath, nil
		}
	}
	return savePath, nil
}

func (cli ftpClient) Readdir(ns string) (string, error) {
	savePath, getPathErr := cli.GetPath(ns)
	if getPathErr != nil {
		return "", getPathErr
	}
	fmt.Println(savePath)
	fmt.Println(savePath)
	fs, err := cli.Client.ReadDir(savePath)
	if err != nil {
		fmt.Println("read dir err: ", err)
		return "", err
	}

	if len(fs) > 0 {
		for _, f := range fs {
			fmt.Println("--------- \n", f.Name())
		}
	}
	return savePath, nil
}

func (cli ftpClient) FtpDel(ns string) error {
	savePath, getPathErr := cli.GetPath(ns)
	if getPathErr != nil {
		return getPathErr
	}
	err := cli.Client.Delete(savePath)
	if err != nil {
		fmt.Println("del path err: ", err)
		return err
	}

	return nil
}
