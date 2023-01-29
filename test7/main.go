package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//fw()
	fr()
}
func fw() {
	filePath := "C:Users/asus/Desktop/a.txt"
	//打开这个地址
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	str := "回八维写作业----->!\r\n" // \r\n 表示换行
	//构造一个写对象
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}

func fr() {
	file, _ := os.Open("C:Users/asus/Desktop/a.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束...")
}
