package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	ff, err := os.Open("./hanying.dat")
	if err != nil {
		log.Fatalln("文件打开错误异常", err.Error())
		//panic("文件打开错误异常")
	}
	defer ff.Close()

	reader := bufio.NewReader(ff)
	// i := 0
	linerstr, err := reader.ReadBytes('\t')
	//头部
	fmt.Println(string(linerstr[16:]))
	//k
	// for {
	// 	linerstr, err := reader.ReadBytes('\t')
	// 	if err != nil {
	// 		break
	// 	}
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(linerstr[10:])
	// 	i++
	// }

}
