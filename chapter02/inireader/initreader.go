package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(getValue("config.ini", "core", "filemode"))
	//获取url
	fmt.Println(getValue("config.ini", "remote \"origin\"", "url"))
}

//getValue 重文件中读取段中的值
func getValue(filepath, selection, expectkey string) (string, error) {
	//读取文件
	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}
	//在函数结束是、时，关闭文件
	defer file.Close()

	//使用读取器读取文件
	reader := bufio.NewReader(file)

	var selectionName string

	for {
		linerstr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//去除两边空白字符串
		linerstr = strings.TrimSpace(linerstr)
		//忽略空白行
		if linerstr == "" {
			continue
		}
		//忽略注释
		if linerstr[0] == ';' {
			continue
		}
		//行首和尾巴分别是方括号，说明是段标记，说明是段标记的起止符
		if linerstr[0] == '[' && linerstr[len(linerstr)-1] == ']' {
			//将段名取出
			selectionName = linerstr[1 : len(linerstr)-1]
		} else if selectionName == selection { //这个段是否是希望读取的值
			//切开等号分隔建值对
			pair := strings.Split(linerstr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if expectkey == key {
					return strings.TrimSpace(pair[1]), nil
				}

			}
		}

	}
	return "", nil
}
