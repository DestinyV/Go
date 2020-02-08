package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	jsonFile, err := os.Open("./information.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer jsonFile.Close()

	// tempContainer := make([]byte, 256)

	// read file
	// for {
	// 	n, err := jsonFile.Read(tempContainer)
	// 	if err != nil {
	// 		fmt.Println("read file error")
	// 		return
	// 	}
	// 	fmt.Println(n)
	// 	fmt.Printf("read content is :%v\n", string(tempContainer))
	// 	if n < 256 {
	// 		return
	// 	}
	// }

	// readByBufio(jsonFile)
	// readByIoutil("./information.json")
	// writeByOpenFile("./logInfor.txt", "hello\nworld")
	// readByIoutil("./logInfor.txt")
	// writeByBufio("./logInfor.txt", "hello\nworld\nBufio")
	writeByIOUtil("./logInfor.txt", "hello\nworld\nioutil")
}

// bufio - read only on line
func readByBufio(f *os.File) {
	reader := bufio.NewReader(f)
	for {
		jsonLine, err := reader.ReadString('\r') // 字符-换行符
		if err == io.EOF {
			fmt.Println("this line read over")
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(jsonLine)
	}
}

// ioutil - read all file
func readByIoutil(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。

// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
// 	...
// }
// 模式	含义
// os.O_WRONLY	只写
// os.O_CREATE	创建文件
// os.O_RDONLY	只读
// os.O_RDWR	读写
// os.O_TRUNC	清空
// os.O_APPEND	追加
// perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

// os.OpenFile write into file
func writeByOpenFile(fileName string, content string) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// logFile.Write([]byte("test byte"))
	len, err := logFile.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("length of writted content is:%d\n", len)
	defer logFile.Close()
}

// bufio
func writeByBufio(fileName string, logContent string) {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()
	writer := bufio.NewWriter(logFile)
	writer.WriteString(logContent) // 写入缓存
	writer.Flush()                 // 将缓存中的信息写入文件
}

// ioutil
func writeByIOUtil(fileName string, logContent string) {
	err := ioutil.WriteFile(fileName, []byte(logContent), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
