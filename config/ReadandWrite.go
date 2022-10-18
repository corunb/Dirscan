package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ReadFile 读取文件列表返回数组
func ReadFile(fileName string) []string {
	list, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", fileName, err)
		os.Exit(0)
	}
	//defer list.Close()

	var lists []string
	scanner := bufio.NewScanner(list)
	scanner.Split(bufio.ScanLines)
	var text string
	for scanner.Scan() {
		text = strings.TrimSpace(scanner.Text())
		lists = append(lists, text)
	}
	return lists
}

// Write 写文件
func Write(output string,url string) {
	//如果没有test.txt这个文件那么就创建，并且对这个文件只进行写和追加内容。
	if Outfile == "" {
		times := time.Now().Format("2006_01_02_")
		fileName := "./" + times + Urladdress(url)+".txt"
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("文件错误,错误为:%v\n", err)
			return
		}
		//defer file.Close()
		str := []byte(output + "\n")
		_,_ = file.Write(str) //将str字符串的内容写到文件中，强制转换为byte，因为Write接收的是byte。
	}else {
		file, err := os.OpenFile(Outfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("文件错误,错误为:%v\n", err)
			return
		}
		//defer file.Close()
		str := []byte(output + "\n")
		_,_ = file.Write(str) //将str字符串的内容写到文件中，强制转换为byte，因为Write接收的是byte。
	}



}