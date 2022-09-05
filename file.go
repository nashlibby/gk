package gk

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// 判断文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// IsNotExist来判断，是不是不存在的错误
	if os.IsNotExist(err) { // 如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err // 如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

// 判断字符串是否在文件中
func StringExistsInFile(filePath string, targetString string) (bool, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	for buf.Scan() {
		if strings.Contains(buf.Text(), strings.TrimSpace(targetString)) {
			return true, nil
		}
	}
	return false, nil
}

// 获取文件最后一行内容
func GetLastLineInFile(filePath string) (string, error) {
	var lastLine string
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	for buf.Scan() {
		lastLine = buf.Text()
	}
	return lastLine, nil
}

// 文件中插入一行内容
func InsertOneLineToFile(filePath string, content string, targetString string, position ...string) error {
	insertPosition := "after"
	if len(position) > 0 {
		insertPosition = position[0]
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	lastLine, err := GetLastLineInFile(filePath)
	if err != nil {
		return err
	}

	if lastLine != "" {
		n, _ := file.Seek(0, io.SeekEnd)
		_, err = file.WriteAt([]byte("\n"), n)
		if err != nil {
			return err
		}
		file.Seek(0, 0)
	}

	buf := bufio.NewReader(file)
	var chunks []byte
	// 按行读取文件内容
	for {
		// 读取一行
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		if insertPosition == "after" {
			chunks = append(chunks, line...)
			if strings.TrimSpace(strings.Trim(line, "\n")) == targetString {
				chunks = append(chunks, content+"\n"...)
			}
		}

		if insertPosition == "before" {
			if strings.TrimSpace(strings.Trim(line, "\n")) == targetString {
				chunks = append(chunks, content+"\n"...)
			}
			chunks = append(chunks, line...)
		}
	}

	// 写入文件
	file.Seek(0, 0)
	file.Truncate(0)
	_, err = file.Write(chunks)
	if err != nil {
		return err
	}
	return nil
}
