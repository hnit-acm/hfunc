package main

import (
	"embed"
	"fmt"
	"github.com/hnit-acm/hfunc/hbasic"
	"github.com/hnit-acm/hfunc/hutils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template
var templateFiles embed.FS

func newService(name, port hbasic.String) bool {
	fileList, _ := ioutil.ReadDir("./")
	for _, fileInfo := range fileList {
		// 如果存在模板文件
		if fileInfo.IsDir() && fileInfo.Name() == "template" {
			_, err := os.Open(name.GetNative())
			if err == nil {
				fmt.Println("服务已存在")
				return false
			}
			copyDir(fileInfo.Name(), name.GetNative(), map[string]string{
				"serviceName": strings.Title(name.GetNative()),
				"port":        port.GetNative(),
			})
			return true
		}
	}
	copyDir("template", name.GetNative(), map[string]string{
		"serviceName": strings.Title(name.GetNative()),
		"port":        port.GetNative(),
	})
	return true
}

func copyDir(src, dest string, params map[string]string) (err error) {
	fileList, _ := templateFiles.ReadDir(src)
	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileList {
		// 如果是目录
		if info.IsDir() {
			if info.Name() == ".git" {
				continue
			}
			// 新建目录
			err = os.MkdirAll(filepath.Join(dest, info.Name()), os.ModePerm)
			if err != nil {
				return
			}
			copyDir(filepath.Join(src, info.Name()), filepath.Join(dest, info.Name()), params)
			continue
		}
		t, err := template.New(info.Name()).Funcs(template.FuncMap{
			"toSnakeString": hutils.StringToSnakeCasedString,
		}).ParseFS(templateFiles, filepath.Join(src, info.Name()))
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
		err = t.Execute(f, params)
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
