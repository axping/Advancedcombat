package main

import "fmt"

const (
	ArrowWeapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

//自定义一个数据类型
type ChipType int

//使用新的数据类型创建一个新的常量，值使用iota
const (
	None ChipType = iota
	CPU
	GPU
	NPU
)

//转换未字符类型
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	case NPU:
		return "NPU"
	}
	return "N/A"
}

func main() {
	//出书CPU的值，并以整型格式输出
	fmt.Printf("%s %d", CPU.String(), CPU)
}
