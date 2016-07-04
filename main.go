package main

import (
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	dirPath    = flag.String("dir_path", "testDir", "The dir path for traverse.")
	resultPath = flag.String("result_path", "result", "The file path for save result.")
	ignorePath = flag.String("ignore_path", "", "The filepath or dirpath for ignore,use commas to separate multiple parameters.")
)

func main() {
	flag.Parse()
	err := TraverseDir(*dirPath, *resultPath, *ignorePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TraverseDir %s success, the result save at %s!\n", *dirPath, *resultPath)
}

// 对文件名进行sha1加密
func Sha1(file string) string {
	s := sha1.New()
	io.WriteString(s, file)
	return fmt.Sprintf("%x", s.Sum(nil))
}

// 判断路径是否存在
func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

// 对目录进行遍历
func TraverseDir(dirPath, resultPath, ignorePath string) error {
	// dirPath有效性检测
	if !PathIsExist(dirPath) {
		panic(errors.New("The dir path not exist!"))
	}
	// result文件检测
	var result *os.File
	var err error
	// 文件不存在创建文件
	if !PathIsExist(resultPath) {
		result, err = os.OpenFile(resultPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
	}
	// 文件存在打开文件
	result, err = os.OpenFile(resultPath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	// ignore文件
	ignorePaths := strings.Split(ignorePath, ",")

	return filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		// 忽略指定目录或者文件
		for _, v := range ignorePaths {
			if v != "" {
				r := regexp.MustCompile(v)
				if r.MatchString(path) {
					fmt.Println("ignore file :" + path)
					return nil
				}
			}
		}
		if !f.IsDir() {
			// 读取文件内容
			fi, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fi.Close()
			content, err := ioutil.ReadAll(fi)
			if err != nil {
				return err
			}
			// 将文件信息写入result文件
			_, err = io.WriteString(result, path+","+Sha1(string(content))+","+fmt.Sprintf("%d", f.Size())+"B\n")
			if err != nil {
				return err
			}
		}
		return nil
	})
}
