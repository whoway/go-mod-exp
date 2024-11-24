package main

import (
	two "exp1/iamsorry"
	"exp1/one"
	"fmt"
)

func main() {
	fmt.Println("main function")
	one.Demo()
	two.Demo() // 坑：注意，虽然你导入的是iamsoryy这个文件夹，但是这里使用的是two这个包
}
