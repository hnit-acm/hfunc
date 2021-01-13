package main

import (
	"fmt"
	"github.com/hnit-acm/hfunc/basic"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func newService(name basic.String) bool {
	fileList, _ := ioutil.ReadDir("./")
	for _, fileInfo := range fileList {
		// 如果存在模板文件
		if fileInfo.IsDir() && fileInfo.Name() == "template" {
			_, err := os.Open(name.GetNative())
			if err == nil {
				fmt.Println("服务已存在")
				return false
			}
			copyDir(fileInfo.Name(), name.GetNative(), name.GetNative())
			return true
		}
	}
	fmt.Println("不存在模板文件,正在下载模板文件")
	if fetchTemplate() != nil {
		fmt.Println("下载模板文件失败")
		return false
	}
	fmt.Println("下载模板文件成功")
	copyDir("template", name.GetNative(), name.GetNative())
	return true
}

func copyDir(src, dest, serviceName string) (err error) {
	fileList, _ := ioutil.ReadDir(src)
	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileList {
		// 如果是目录
		if info.IsDir() {
			// 新建目录
			err = os.MkdirAll(filepath.Join(dest, info.Name()), os.ModePerm)
			if err != nil {
				return
			}
			copyDir(filepath.Join(src, info.Name()), filepath.Join(dest, info.Name()), serviceName)
			continue
		}
		// 文件
		//data, _ := ioutil.ReadFile(filepath.Join(src, info.Name()))
		t, err := template.ParseFiles(filepath.Join(src, info.Name()))
		if err != nil {
			fmt.Println(err)
			return err
		}

		f, err := os.OpenFile(filepath.Join(dest, strings.TrimSuffix(info.Name(), ".ht")), os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return err
		}
		err = t.Execute(f, map[string]string{"service_name": strings.Title(serviceName)})
		if err != nil {
			f.Close()
			fmt.Println(err)
			continue
		}
		f.Close()
		continue
	}
	return err
}

func fetchTemplate() error {
	e := exec.Command("git", "clone", "https://github.com/hnit-acm/template")
	fmt.Println(e.String())
	err := e.Run()
	if err != nil {
		fmt.Println("下载模板文件失败")
		return err
	}
	return nil
}
