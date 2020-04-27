package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func mkdirFolder(path string) {
	// 创建文件夹
	folderName := "output_" + time.Now().Format("2006-01-02")
	os.Mkdir(path+"/"+folderName, os.ModePerm)
	fmt.Printf("%s 文件夹创建成功 \n", folderName)
	GetAllFile(path, folderName)
}

func GetAllFile(pathname string, folderName string) string {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return "read dir fail"
	}
	if len(rd) != 0 {
		for _, fi := range rd {
			if fi.IsDir() {
				fullDir := pathname + "/" + fi.Name()
				// 遍历文件夹
				GetAllFile(fullDir, folderName)
				if err != nil {
					fmt.Println("read dir fail:", err)
					return "read dir fail"
				}
			} else {
				fullName := pathname + "/" + fi.Name()
				isVueFile, _ := regexp.Match(`.vue$`, []byte(fullName)) // 只选择 .vue 文件
				curPath, _ := os.Getwd()
				if isVueFile {
					outFilePath := curPath + "/" + folderName + "/" + fi.Name()
					redFile(fullName, strings.Replace(outFilePath, ".vue", ".json", -1))
				}
			}
		}
	}
	return "ok"
}

// 读取文件
func redFile(path string, outPath string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	regExce := regexp.MustCompile(`<i18n>([\S\s]*)<\/i18n>`)
	s := regExce.FindStringSubmatch(string(file))
	if len(s) != 0 {
		// 写入文件
		outFileErr := ioutil.WriteFile(outPath, []byte(s[1]), 0666)
		if outFileErr != nil {
			fmt.Println("write fail", outFileErr)
		}
		fmt.Printf("%s write success \n", outPath)
	}
	fmt.Println("全部成功")
}

func main() {
	// 获取当前路径
	curPath, _ := os.Getwd()
	mkdirFolder(curPath)
}
