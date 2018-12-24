package main

import (
	"fmt"
	"github.com/garenwen/godep-feature-dep/types"
)

func main() {
	user := new(types.User)
	user.Name = "Garen wen"
	fmt.Printf("user is %v",user)
	fmt.Println("Hi ")
}
