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

func main() {
	// dirの作成
	desiredDir := os.Args[1]
	// 現在のcurrent dirを取得
	currentDir, _ := os.Getwd()

	createdDirPath, ok := makeDir(desiredDir, currentDir)
	fmt.Printf("dirを作成しました %s", createdDirPath)
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
	isString := interface{}(dirName)

	fullPath := currentDir +  "/" + dirName
	fmt.Println(fullPath)
	//すでにdirがあるか確認
	_, err := os.Stat(fullPath)
	if !os.IsNotExist(err) {
		// 存在していればnilを返し処理を継続
		return dirName, nil
	}

	if _, ok := isString.(string); !ok {
		log.Fatal("第一引数は文字列を入力してください")
	}

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

//  copyに変更したい。
func makeFile(fileName, path string) error {
	filepath := path + "/" + fileName
	io.Copy()
		_, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
	return nil
} 
