package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
)

//go:embed files/hello.txt
var embedHello string

func fileHello() string {
	b, err := ioutil.ReadFile("files/hello.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	fmt.Printf("embedHello:\t%q\n", embedHello)
	fmt.Printf("fileHello:\t%q\n", fileHello())
}
