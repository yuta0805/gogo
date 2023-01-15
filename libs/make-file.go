package libs

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

const (
	OriginDir = "origin"
)

type makeFileError struct {
	s string
	Err error
}

func (m *makeFileError) Error() string {
	return fmt.Sprintf("message: %s %s\n", m.s, m.Err.Error())
}

func MakeDir(dirName, currentDir string) (string, error) {
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
		os.Exit(1)
	}

	//空文字かどうか
	if dirName == "" {
		log.Fatal("作成するディレクトリ名を入力してください")
		os.Exit(1)
	}

	err = os.Mkdir(fullPath, 0755)
	return fullPath, err
}

func ReadOriginFile(dirName string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal("empty something in dir")
		os.Exit(1)
	}

	return files, nil
}

func MakeFile(fileName, path string) error {
	filePath := path + "/" + fileName

	_, isFile := os.Open(filePath)
	if os.IsExist(isFile) {
		return nil
	}

	writer, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	reader, err := os.Open(OriginDir + "/" + fileName)
	
	_, ok := io.Copy(writer, reader)
	if ok != nil {
		log.Fatal(err)
	}

	return nil
} 
