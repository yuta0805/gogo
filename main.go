package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

const (
	originDir = "origin"
)


type File string

func (f File) Writer (p []byte) (n int, err error) {
	return len(f), nil
}

func main() {
	// dirの作成
	desiredDir := os.Args[1]
	// 現在のcurrent dirを取得
	currentDir, _ := os.Getwd()

	createdDirPath, ok := makeDir(desiredDir, currentDir)
	fmt.Printf("dirを作成しました %s\n", createdDirPath)
	if ok != nil {
		log.Fatal(ok)
	}

	//origin dir配下のterraformファイル情報を取得
	files, isFiles := readOriginFile(originDir)
	if isFiles != nil {
		log.Fatal("ファイル読み込みに失敗しました")
	}

	// fileの生成
	for _, file := range files {
		makeFile(file.Name(), createdDirPath)
	}
}

func makeDir(dirName, currentDir string) (string, error) {
	fullPath := currentDir +  "/" + dirName
	//すでにdirがあるか確認
	_, err := os.Stat(fullPath)
	if !os.IsNotExist(err) {
		// 存在していればnilを返し処理を継続
		return dirName, nil
	}

	//型のチェック
	isString := interface{}(dirName)
	if _, ok := isString.(string); !ok {
		log.Fatal("第一引数は文字列を入力してください")
	}

	//空文字かどうか
	if dirName == "" {
		log.Fatal("作成するディレクトリ名を入力してください")
	}

	err = os.Mkdir(fullPath, 0755)
	return fullPath, err
}

func readOriginFile(dirName string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal("empty something in dir")
	}

	return files, nil
}

func makeFile(fileName, path string) error {
	filePath := path + "/" + fileName

	_, isFile := os.Open(filePath)
	if os.IsExist(isFile) {
		return nil
	}

	writer, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	reader, err := os.Open(originDir + "/" + fileName)
	
	_, ok := io.Copy(writer, reader)
	if ok != nil {
		log.Fatal(err)
	}

	return nil
} 
