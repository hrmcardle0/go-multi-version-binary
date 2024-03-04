package main

import (
	"fmt"
	"github.com/hrmcardle0/go-multi-version-module/v2/mymodule"
)

func main() {
	fmt.Println("Inside main()")
	ret := mymodulelongname.ModuleFuncv2()
	fmt.Println(ret)
}
