package main

import (
	"fmt"
	"log"
	"os"

	"gogo/libs"
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

	createdDirPath, ok := libs.MakeDir(desiredDir, currentDir)
	fmt.Printf("dirを作成しました %s\n", createdDirPath)
	if ok != nil {
		log.Fatal(ok)
	}

	//origin dir配下のterraformファイル情報を取得
	files, isFiles := libs.ReadOriginFile(libs.OriginDir)
	if isFiles != nil {
		isFiles.Error()
		log.Fatal("ファイル読み込みに失敗しました")
	}

	// fileの生成
	for _, file := range files {
		libs.MakeFile(file.Name(), createdDirPath)
	}
}
