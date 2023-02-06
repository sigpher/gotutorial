package main

import (
	"fmt"
	"hello/funcpack"
)

func main() {
	ret := funcpack.Add(10, 20)
	fmt.Println(ret)

	txtFunc := funcpack.MakeSuffix("txt")
	fmt.Println(txtFunc("myAsset"))

}
