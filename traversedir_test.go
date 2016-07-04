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
		"testDir/.git,5222aba97615966f4a4cdab21e0f1df7d636c468,12B\n",
		"testDir/.gitignore,354393ef86c1a16cbd2e8a8c778bedc833fc18c8,19B\n",
		"testDir/README,f78a71af8bbf8cc2f6f313549d4da14bd3771359,6B\n",
		"testDir/test/.vscode,a94a8fe5ccb19ba61c4c0873d391e987982fbbd3,4B\n",
		"testDir/test/111,6216f8a75fd5bb3d5f22b6f9958cdede3fc086c2,3B\n",
		"testDir/test/222,1c6637a8f2e1f75e06ff9984894d6bd16a3a36a9,3B\n",
		"testDir/test1,b444ac06613fc8d63795be9ad0beaf55011936ac,5B\n",
		"testDir/test2,109f4b3c50d7b0df729d299bc6f8e9ef9066971f,5B\n",
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

func TestIgnore1(t *testing.T) {
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
		"testDir/README,f78a71af8bbf8cc2f6f313549d4da14bd3771359,6B\n",
		"testDir/test/111,6216f8a75fd5bb3d5f22b6f9958cdede3fc086c2,3B\n",
		"testDir/test/222,1c6637a8f2e1f75e06ff9984894d6bd16a3a36a9,3B\n",
		"testDir/test1,b444ac06613fc8d63795be9ad0beaf55011936ac,5B\n",
		"testDir/test2,109f4b3c50d7b0df729d299bc6f8e9ef9066971f,5B\n",
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

func TestIgnore2(t *testing.T) {
	flag.Parse()
	if !PathIsExist(*dirPath) {
		panic(errors.New("The dir path not exist!"))
	}

	// 删除原有文件
	os.Remove(*resultPath)

	err := TraverseDir(*dirPath, *resultPath, ".git")
	if err != nil {
		t.Error(err)
	}

	// 对比结果文件
	exceptResult := []string{
		"testDir/README,f78a71af8bbf8cc2f6f313549d4da14bd3771359,6B\n",
		"testDir/test/.vscode,a94a8fe5ccb19ba61c4c0873d391e987982fbbd3,4B\n",
		"testDir/test/111,6216f8a75fd5bb3d5f22b6f9958cdede3fc086c2,3B\n",
		"testDir/test/222,1c6637a8f2e1f75e06ff9984894d6bd16a3a36a9,3B\n",
		"testDir/test1,b444ac06613fc8d63795be9ad0beaf55011936ac,5B\n",
		"testDir/test2,109f4b3c50d7b0df729d299bc6f8e9ef9066971f,5B\n",
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
