package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	flag.Parse()
	if !PathIsExist(*dirPath) {
		panic(errors.New("The dir path not exist!"))
	}

	// 删除原有文件
	os.Remove(*resultPath)

	err := TraverseDir(*dirPath, *resultPath, *ignorePath)
	if err != nil {
		t.Error(err)
	}

	// 对比结果文件
	exceptResult := []string{
		"testDir/README,437b3abd82653ac96d56cc06adaf30f559c2e9c2,6B\n",
		"testDir/test/.vscode,983938d6d9aea270b2fc8a01a869fb17239066c7,4B\n",
		"testDir/test/111,de198a14f4203caacb8dbc0bdc26d000fff4e890,3B\n",
		"testDir/test/222,f3dad947bcc2b2b5a96c77e6ec67af67f01f55c9,3B\n",
		"testDir/test1,f93465497ab87166ab66f07bd31dec781a486039,5B\n",
		"testDir/test2,c67cd5b5c52f156ef11540f589d0acb108932c4b,5B\n",
	}

	// 读取result文件
	f, err := os.Open(*resultPath)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for i := 0; i < 5; i++ {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		assert.Equal(t, exceptResult[i], line)
	}
}

func TestIgnore(t *testing.T) {
	flag.Parse()
	if !PathIsExist(*dirPath) {
		panic(errors.New("The dir path not exist!"))
	}

	// 删除原有文件
	os.Remove(*resultPath)

	err := TraverseDir(*dirPath, *resultPath, "[.]")
	if err != nil {
		t.Error(err)
	}

	// 对比结果文件
	exceptResult := []string{
		"testDir/README,437b3abd82653ac96d56cc06adaf30f559c2e9c2,6B\n",
		"testDir/test/111,de198a14f4203caacb8dbc0bdc26d000fff4e890,3B\n",
		"testDir/test/222,f3dad947bcc2b2b5a96c77e6ec67af67f01f55c9,3B\n",
		"testDir/test1,f93465497ab87166ab66f07bd31dec781a486039,5B\n",
		"testDir/test2,c67cd5b5c52f156ef11540f589d0acb108932c4b,5B\n",
	}

	// 读取result文件
	f, err := os.Open(*resultPath)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for i := 0; i < 5; i++ {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		assert.Equal(t, exceptResult[i], line)
	}

}
