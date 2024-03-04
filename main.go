package main

import (
	"fmt"
	"github.com/hrmcardle0/go-multi-version-module/mymodule"
)

func main() {
	fmt.Println("Inside main()")
	ret := mymodule.ModuleFunc()
	fmt.Println(ret)
}
